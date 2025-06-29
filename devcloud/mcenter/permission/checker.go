// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package permission

import (
	"context"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
)

func init() {
	ioc.Config().Registry(&Checker{})
}

func Auth(v bool) (string, bool) {
	return endpoint.MetaRequiredAuthKey, v
}

func Permission(v bool) (string, bool) {
	return endpoint.MetaRequiredPermKey, v
}

func Resource(v string) (string, string) {
	return endpoint.MetaResourceKey, v
}

func Action(v string) (string, string) {
	return endpoint.MetaActionKey, v
}

type Checker struct {
	ioc.ObjectImpl
	log *zerolog.Logger

	token  token.Service
	policy policy.Service
}

func (c *Checker) Name() string {
	return "permission_checker"
}

// Priority 初始化的优先级,由大到小,默认是 0,此时必须比框架小一点,框架是 899
// 框架的 Init 函数调用完成后立马调用这个对象的 Init 函数,实现了全局中间件的功能
func (c *Checker) Priority() int {
	return gorestful.Priority() - 1
}

func (c *Checker) Init() error {
	c.log = log.Sub(c.Name())
	c.token = token.GetService()
	c.policy = policy.GetService()

	gorestful.RootRouter().Filter(c.Check)
	return nil
}

func (c *Checker) Check(r *restful.Request, w *restful.Response, chain *restful.FilterChain) {
	// 知道用户当前访问的是哪个路由
	// SelectedRoute 返回当前 URL 适配哪个路由
	route := endpoint.NewEntryFromRestRouteReader(r.SelectedRoute())
	if route.RequiredAuth {
		// 需要鉴权
		tk, err := c.CheckToken(r)
		if err != nil {
			response.Failed(w, err)
			return
		}

		// 校验权限
		if err := c.CheckPolicy(r, tk, route); err != nil {
			response.Failed(w, err)
			return
		}
	}

	// 请求处理
	chain.ProcessFilter(r, w)
	// 请求处理后
}

func (c *Checker) CheckToken(r *restful.Request) (*token.Token, error) {
	// 从请求头或 cookie 里获取 token 信息
	v := token.GetAccessTokenFromHTTP(r.Request)
	if v == "" {
		return nil, exception.NewUnauthorized("请先登录")
	}

	// 调用 token 服务验证 token 是否合法
	tk, err := c.token.ValidateToken(r.Request.Context(), token.NewValidateTokenRequest(v))
	if err != nil {
		return nil, err
	}

	// 如果校验成功,需要把用户的身份信息,保存到请求的上下文中,方便后续逻辑获取
	ctx := context.WithValue(r.Request.Context(), token.CtxTokenKey, tk)

	// 把新的请求对象,重新赋值给当前的请求
	r.Request = r.Request.WithContext(ctx)
	return tk, nil
}

// CheckPolicy 验证策略
// 获取当前访问的路由
func (c *Checker) CheckPolicy(r *restful.Request, tk *token.Token, route *endpoint.RouteEntry) error {
	// 判断用户是否是超级管理员,如果是的话,就不用继续了
	if tk.IsAdmin {
		return nil
	}

	if route.HasRequiredRole() {
		set, err := c.policy.QueryPolicy(r.Request.Context(),
			policy.NewQueryPolicyRequest().
				SetNamespaceId(tk.NamespaceId).
				SetUserId(tk.UserId).
				SetExpired(false).
				SetEnabled(true).
				SetWithRole(true),
		)
		if err != nil {
			return exception.NewInternalServerError("%s", err.Error())
		}
		hasPerm := false
		for i := range set.Items {
			p := set.Items[i]
			if route.IsRequireRole(p.Role.Name) {
				hasPerm = true
				break
			}
		}

		if !hasPerm {
			return exception.NewPermissionDeny("无权限访问")
		}
	}

	// api 权限校验
	if route.RequiredPerm {
		validateReq := policy.NewValidateEndpointPermissionRequest()
		validateReq.UserId = tk.UserId
		validateReq.NamespaceId = tk.NamespaceId
		validateReq.Service = application.Get().GetAppName()
		validateReq.Method = route.Method
		validateReq.Path = route.Path
		resp, err := c.policy.ValidateEndpointPermission(r.Request.Context(), validateReq)
		if err != nil {
			return exception.NewInternalServerError("%s", err.Error())
		}

		if !resp.HasPermission {
			return exception.NewPermissionDeny("无权限访问")
		}
	}
	return nil
}

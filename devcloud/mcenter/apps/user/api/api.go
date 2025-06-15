// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/sword-demon/go18/devcloud/audit/comptroller"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"github.com/sword-demon/go18/devcloud/mcenter/permission"
)

func init() {
	// 注册用户模块服务
	ioc.Api().Registry(&UserRestfulApiHandler{})
}

type UserRestfulApiHandler struct {
	ioc.ObjectImpl

	svc user.Service // 依赖 user 服务实现
}

func (h *UserRestfulApiHandler) Name() string {
	return user.AppName
}

func (h *UserRestfulApiHandler) Init() error {
	h.svc = user.GetService() // 获取服务

	tags := []string{"用户接口"} // 文档的 tag
	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("").To(h.CreateUser).
		Doc("创建用户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(user.User{}).
		Returns(200, "OK", user.User{}))

	ws.Route(ws.GET("").To(h.QueryUser).
		Doc("查询用户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 表示这个接口需要进行鉴权
		Metadata(permission.Auth(true)).
		Metadata(permission.Action("list")).
		Metadata(permission.Resource("user")).
		Metadata(comptroller.Enable(true)). // 开启审计
		Param(ws.QueryParameter("user_id", "用户ID数组,案例 user_id=1&user_id=2").DataType("string")).
		Param(ws.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(ws.QueryParameter("page_number", "分页页码").DataType("integer")).
		Writes(QuerySet{}).
		Returns(200, "OK", QuerySet{}))

	return nil
}

// QuerySet go-restful 的文档模式不支持泛型,所以这里额外进行定义
type QuerySet struct {
	Total int64       `json:"total"`
	Items []user.User `json:"items"`
}

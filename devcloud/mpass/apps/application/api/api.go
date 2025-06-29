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
	"github.com/sword-demon/go18/devcloud/mcenter/permission"
	"github.com/sword-demon/go18/devcloud/mpass/apps/application"
)

func init() {
	// 注册用户模块服务
	ioc.Api().Registry(&ApplicationRestfulApiHandler{})
}

type ApplicationRestfulApiHandler struct {
	ioc.ObjectImpl

	svc application.Service // 依赖 user 服务实现
}

func (h *ApplicationRestfulApiHandler) Name() string {
	return application.AppName
}

func (h *ApplicationRestfulApiHandler) Init() error {
	h.svc = application.GetService() // 获取服务

	tags := []string{"应用管理"} // 文档的 tag
	ws := gorestful.ObjectRouter(h)

	ws.Route(ws.GET("").To(h.QueryApplication).
		Doc("查询应用列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 表示这个接口需要进行鉴权
		Metadata(permission.Auth(true)).
		Metadata(permission.Action("list")).
		Metadata(permission.Resource("application")).
		Metadata(comptroller.Enable(true)). // 开启审计
		Param(ws.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(ws.QueryParameter("page_number", "分页页码").DataType("integer")).
		Writes(QuerySet{}).
		Returns(200, "OK", QuerySet{}))

	return nil
}

// QuerySet go-restful 的文档模式不支持泛型,所以这里额外进行定义
type QuerySet struct {
	Total int64                     `json:"total"`
	Items []application.Application `json:"items"`
}

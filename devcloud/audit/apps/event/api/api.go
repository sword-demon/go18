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
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
)

func init() {
	ioc.Api().Registry(&EventRestfulApiHandler{})
}

type EventRestfulApiHandler struct {
	ioc.ObjectImpl

	event event.Service
}

func (h *EventRestfulApiHandler) Name() string {
	return event.AppName
}

func (h *EventRestfulApiHandler) Init() error {
	h.event = event.GetService() // 获取事件服务

	tags := []string{"审计接口"} // 文档的 tag
	ws := gorestful.ObjectRouter(h)

	ws.Route(ws.GET("").To(h.QueryEventApi).
		Doc("查询审计日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		//Metadata(comptroller.Enable(true)).
		// 表示这个接口需要进行鉴权
		Param(ws.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(ws.QueryParameter("page_number", "分页页码").DataType("integer")).
		Writes(QuerySet{}).
		Returns(200, "OK", QuerySet{}))

	return nil
}

type QuerySet struct {
	Total int64         `json:"total"`
	Items []event.Event `json:"items"`
}

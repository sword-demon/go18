// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package application

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
)

const (
	AppName = "application"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	CreateApplication(context.Context, *CreateApplicationRequest) (*Application, error)           // 创建应用
	QueryApplication(context.Context, *QueryApplicationRequest) (*types.Set[*Application], error) // 查询应用
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error)           // 更新应用
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*Application, error)           // 删除应用
	DescribeApplication(context.Context, *DescribeApplicationRequest) (*Application, error)       // 查询应用详情
}

type QueryApplicationRequest struct {
	policy.ResourceScope
	QueryApplicationRequestSpec
}

func NewQueryApplicationRequest() *QueryApplicationRequest {
	return &QueryApplicationRequest{
		QueryApplicationRequestSpec: QueryApplicationRequestSpec{
			PageRequest: request.NewDefaultPageRequest(),
		},
	}
}

type QueryApplicationRequestSpec struct {
	*request.PageRequest
	Id    string `json:"id" bson:"_id"`      // 应用ID
	Name  string `json:"name" bson:"name"`   // 应用名称
	Ready *bool  `json:"ready" bson:"ready"` // 是否就绪
}

type UpdateApplicationRequest struct {
	UpdateBy string `json:"update_by" bson:"update_by"` // 更新人
	DescribeApplicationRequest
	CreateApplicationSpec
}

type DescribeApplicationRequest struct {
	policy.ResourceScope
	Id string `json:"id" bson:"_id"` // 应用ID
}

type DeleteApplicationRequest struct {
	DescribeApplicationRequest
}

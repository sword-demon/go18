// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package label

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

const (
	AppName = "label"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	CreateLabel(context.Context, *CreateLabelRequest) (*Label, error)
	UpdateLabel(context.Context, *UpdateLabelRequest) (*Label, error)
	DeleteLabel(context.Context, *DeleteLabelRequest) (*Label, error)
	DescribeLabel(context.Context, *DescribeLabelRequest) (*Label, error)
	QueryLabel(context.Context, *QueryLabelRequest) (*types.Set[*Label], error)
}

type UpdateLabelRequest struct {
	DescribeLabelRequest
	UpdateBy string           `json:"update_by"` // 更新人
	Spec     *CreateLabelSpec `json:"spec"`      // 标签信息
}

type DeleteLabelRequest struct {
	DescribeLabelRequest
}

type DescribeLabelRequest struct {
	Id string `json:"id"` // 标签 id
}

func NewDescribeLabelRequest() *DescribeLabelRequest {
	return &DescribeLabelRequest{}
}

func (d *DescribeLabelRequest) SetId(Id string) {
	d.Id = Id
}

type QueryLabelRequest struct {
	*request.PageRequest
}

func NewQueryLabelRequest() *QueryLabelRequest {
	return &QueryLabelRequest{
		PageRequest: request.NewDefaultPageRequest(),
	}
}

// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package event

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

var (
	AppName = "event"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	SaveEvent(context.Context, *types.Set[*Event]) error
	QueryEvent(context.Context, *QueryEventRequest) (*types.Set[*Event], error)
}

type QueryEventRequest struct {
	// 分页参数
	*request.PageRequest
}

func NewQueryEventRequest() *QueryEventRequest {
	return &QueryEventRequest{
		PageRequest: request.NewDefaultPageRequest(),
	}
}

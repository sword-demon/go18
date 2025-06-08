// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
)

func init() {
	ioc.Controller().Registry(&EndpointServiceImpl{})
}

var _ endpoint.Service = (*EndpointServiceImpl)(nil)

type EndpointServiceImpl struct {
	ioc.ObjectImpl
}

func (s *EndpointServiceImpl) Name() string {
	return endpoint.AppName
}

func (s *EndpointServiceImpl) Init() error {
	// 自动创建表
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&endpoint.Endpoint{})
		if err != nil {
			return err
		}
	}

	return nil
}

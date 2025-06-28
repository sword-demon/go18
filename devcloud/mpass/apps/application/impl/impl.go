// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mpass/apps/application"
)

func init() {
	ioc.Controller().Registry(&ApplicationServiceImpl{})
}

var _ application.Service = (*ApplicationServiceImpl)(nil)

type ApplicationServiceImpl struct {
	ioc.ObjectImpl
}

func (s *ApplicationServiceImpl) Name() string {
	return application.AppName
}

func (s *ApplicationServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&application.Application{})
		if err != nil {
			return err
		}
	}
	return nil
}

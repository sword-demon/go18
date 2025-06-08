// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
)

func init() {
	ioc.Controller().Registry(&RoleServiceImpl{})
}

var _ role.Service = (*RoleServiceImpl)(nil)

type RoleServiceImpl struct {
	ioc.ObjectImpl
}

func (i *RoleServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&role.Role{}, &role.ApiPermission{}, &role.ViewPermission{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *RoleServiceImpl) Name() string {
	return role.AppName
}

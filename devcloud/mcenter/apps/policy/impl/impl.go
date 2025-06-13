// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
)

func init() {
	ioc.Controller().Registry(&PolicyServiceImpl{})
}

var _ policy.Service = (*PolicyServiceImpl)(nil)

type PolicyServiceImpl struct {
	ioc.ObjectImpl

	namespace namespace.Service
	role      role.Service
}

func (i *PolicyServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&policy.Policy{})
		if err != nil {
			return err
		}
	}

	i.namespace = namespace.GetService()
	i.role = role.GetService()
	return nil
}

func (i *PolicyServiceImpl) Name() string {
	return policy.AppName
}

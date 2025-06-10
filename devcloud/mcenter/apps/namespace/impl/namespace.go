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
)

func init() {
	ioc.Controller().Registry(&NamespaceServiceImpl{})
}

var _ namespace.Service = (*NamespaceServiceImpl)(nil)

type NamespaceServiceImpl struct {
	ioc.ObjectImpl
}

func (i *NamespaceServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&namespace.Namespace{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *NamespaceServiceImpl) Name() string {
	return namespace.AppName
}

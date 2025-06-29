// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/label"
)

func init() {
	ioc.Controller().Registry(&LabelServiceImpl{})
}

var _ label.Service = (*LabelServiceImpl)(nil)

type LabelServiceImpl struct {
	ioc.ObjectImpl
}

func (s *LabelServiceImpl) Name() string {
	return label.AppName
}

func (s *LabelServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&label.Label{})
		if err != nil {
			return err
		}
	}
	return nil
}

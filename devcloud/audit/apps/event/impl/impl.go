// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"github.com/rs/zerolog"
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	ioc.Controller().Registry(&EventServiceImpl{})
}

var _ event.Service = (*EventServiceImpl)(nil)

type EventServiceImpl struct {
	ioc.ObjectImpl

	log *zerolog.Logger

	col *mongo.Collection
}

func (i *EventServiceImpl) Name() string {
	return event.AppName
}

func (i *EventServiceImpl) Init() error {
	// 对象
	i.log = log.Sub(i.Name())

	i.log.Debug().Msgf("database: %s", ioc_mongo.Get().Database)
	// 需要一个集合Collection
	i.col = ioc_mongo.DB().Collection("events")
	return nil
}

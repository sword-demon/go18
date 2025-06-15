// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package comptroller

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	ioc_kafka "github.com/infraboard/mcube/v2/ioc/config/kafka"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

func init() {
	ioc.Config().Registry(&EventSender{Topic: "audit_go18"})
}

type EventSender struct {
	ioc.ObjectImpl

	log *zerolog.Logger

	Topic  string `json:"topic" toml:"topic" yaml:"topic" env:"TOPIC"` // 当前消费者配置的 topic
	writer *kafka.Writer
}

func (sender *EventSender) Name() string {
	return "audit_event_sender"
}

func (sender *EventSender) Init() error {
	sender.log = log.Sub(sender.Name())
	sender.writer = ioc_kafka.Producer(sender.Topic)

	// 注册发送事件 中间件
	gorestful.RootRouter().Filter(sender.SendEvent())
	return nil
}

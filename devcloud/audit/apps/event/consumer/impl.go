// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package consumer

import (
	"context"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_kafka "github.com/infraboard/mcube/v2/ioc/config/kafka"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
)

func init() {
	ioc.Controller().Registry(&consumer{
		GroupId: "audit",
		Topics:  []string{"audit_go18"},
		ctx:     context.Background(),
	})
}

type consumer struct {
	ioc.ObjectImpl

	// 日志
	log *zerolog.Logger
	// kafka 消费者
	reader *kafka.Reader
	// 允许上下文
	ctx context.Context

	// 消费组 id
	GroupId string `json:"group_id" toml:"group_id" yaml:"group_id" env:"GROUP_ID"`
	// 当前这个消费者配置的 topic
	Topics []string `json:"topic" toml:"topic" yaml:"topic" env:"TOPIC"`
}

func (c *consumer) Name() string {
	return "maudit_consumer"
}

func (c *consumer) Init() error {
	c.log = log.Sub(c.Name())
	c.reader = ioc_kafka.ConsumerGroup(c.GroupId, c.Topics)

	go func() {
		err := c.Run(c.ctx)
		if err != nil {
			c.log.Error().Msgf("consumer run error: %s", err)
		}
	}()

	return nil
}

func (c *consumer) Priority() int {
	// 消费者的优先级要比审计日志低
	return event.Priority - 1
}

func (c *consumer) Close(ctx context.Context) error {
	c.ctx.Done()
	return nil
}

// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package consumer

import (
	"context"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
	"io"
)

// Run 读取消息,处理消息
func (c *consumer) Run(ctx context.Context) error {
	// 会阻塞
	for {
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if err == io.EOF {
				c.log.Info().Msg("reader closed")
				return nil
			}
			c.log.Error().Msgf("fetch message error: %s", err)
			continue
		}

		// 处理消息
		e := event.NewEvent()
		c.log.Debug().Msgf("message at topic/partition/offset %v/%v/%v", m.Topic, m.Partition, m.Offset)

		// 发送的数据是 json 格式,接收用的 json,发送也需要 json
		err = e.Load(m.Value)
		if err == nil {
			if err := event.GetService().SaveEvent(ctx, types.NewSet[*event.Event]().Add(e)); err != nil {
				c.log.Error().Msgf("save event error: %s", err)
			}
		}

		// 处理完消息后需要提交该消息,标志为已完成
		if err := c.reader.CommitMessages(ctx, m); err != nil {
			c.log.Error().Msgf("commit message error: %s", err)
		}
	}
}

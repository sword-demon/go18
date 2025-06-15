// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package event

import (
	"encoding/json"
	"github.com/rs/xid"
	"github.com/segmentio/kafka-go"
	"time"
)

// Event 用户操作事件 映射 MongoDB BSON 格式
type Event struct {
	Id           string `json:"id" bson:"_id"` // mongodb 中的 _id 表示的是对象 id
	Who          string `json:"who" bson:"who"`
	Time         int64  `json:"time" bson:"time"`
	Ip           string `json:"ip" bson:"ip"`
	UserAgent    string `json:"user_agent" bson:"user_agent"`
	Service      string `json:"service" bson:"service"`     // Service 服务名称 做了什么操作 服务:资源:动作
	Namespace    string `json:"namespace" bson:"namespace"` // Namespace 命名空间
	ResourceType string `json:"resource_type" bson:"resource_type"`
	Action       string `json:"action" bson:"action"` // 操作类型 <list, get, update, create, delete>

	ResourceId   string `json:"resource_id" bson:"resource_id"` // ResourceId 详情信息
	StatusCode   int    `json:"status_code" bson:"status_code"`
	ErrorMessage string `json:"error_message" bson:"error_message"`

	Label  map[string]string `json:"label" bson:"label"`   // Label 标签
	Extras map[string]string `json:"extras" bson:"extras"` // Extras 扩展信息
}

func NewEvent() *Event {
	return &Event{
		Id:     xid.New().String(),
		Label:  make(map[string]string),
		Extras: make(map[string]string),
		Time:   time.Now().Unix(),
	}
}

func (e *Event) Load(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *Event) ToKafkaMessage() kafka.Message {
	data, _ := json.Marshal(e)
	return kafka.Message{
		Value: data,
	}
}

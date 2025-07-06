// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package resource

import (
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
)

// Resource 资源
// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/bson/#struct-tags inline 的文档描述
type Resource struct {
	Meta   `bson:"inline"`
	Spec   `bson:"inline"`
	Status `bson:"inline"`
}

func NewResource() *Resource {
	return &Resource{
		Meta: Meta{},
		Spec: Spec{
			Tags:  map[string]string{},
			Extra: map[string]string{},
		},
		Status: Status{},
	}
}

func NewResourceSet() *types.Set[*Resource] {
	return types.New[*Resource]()
}

func (r *Resource) String() string {
	return pretty.ToJSON(r)
}

func (r *Resource) TableName() string {
	return "resources"
}

type Meta struct {
	Id        string `bson:"_id" json:"id" validate:"required" gorm:"column:id;primaryKey"` // 唯一 id,直接使用云商的 id
	Domain    string `json:"domain" validate:"required" gorm:"column:domain"`               // 资源所属域
	Namespace string `json:"namespace" validate:"required" gorm:"column:namespace"`         // 资源所属空间
	Env       string `json:"env" gorm:"column:env"`                                         // 资源所属环境
	CreateAt  int64  `json:"create_at" gorm:"column:create_at"`                             // 创建时间
	// 删除时间
	DeleteAt int64 `json:"delete_at" gorm:"column:delete_at"`
	// 删除人
	DeleteBy string `json:"delete_by" gorm:"column:delete_by"`
	// 同步时间
	SyncAt int64 `json:"sync_at" validate:"required" gorm:"column:sync_at"`
	// 同步人
	SyncBy       string `json:"sync_by" gorm:"column:sync_by"`
	CredentialId string `json:"credential_id" gorm:"column:credential_id"` // CredentialId 用于同步的凭证 id
	SerialNumber string `json:"serial_number" gorm:"column:serial_number"` // SerialNumber 序列号
}

// Spec 表单数据,用户申请资源时,需要提供的参数
type Spec struct {
	Vendor        Vendor            `json:"vendor" gorm:"column:vendor"`                         // Vendor 厂商
	ResourceType  Type              `json:"resource_type" gorm:"column:resource_type"`           // ResourceType 资源类型
	Region        string            `json:"region" gorm:"column:region"`                         // Region 区域
	Zone          string            `json:"zone" gorm:"column:zone"`                             // Zone 区域
	Owner         string            `json:"owner" gorm:"column:owner"`                           // Owner 资源所属账号
	Name          string            `json:"name" gorm:"column:name"`                             // Name 名称
	Category      string            `json:"category" gorm:"column:category"`                     // Category 种类
	Type          string            `json:"type" gorm:"column:type"`                             // Type 规格
	Description   string            `json:"description" gorm:"column:description"`               // Description 描述
	ExpireAt      int64             `json:"expire_at" gorm:"column:expire_at"`                   // ExpireAt 过期时间
	UpdateAt      int64             `json:"update_at" gorm:"column:update_at"`                   // UpdateAt 更新时间
	Cpu           int64             `json:"cpu" gorm:"column:cpu"`                               // Cpu 资源占用 cpu 数量
	Gpu           int64             `json:"gpu" gorm:"column:gpu"`                               // Gpu 资源占用的 GPU 数量
	Memory        int64             `json:"memory" gorm:"column:memory"`                         // Memory 资源使用的内存
	SystemStorage int64             `json:"system_storage" gorm:"column:system_storage"`         // SystemStorage 系统盘
	DataStorage   int64             `json:"data_storage" gorm:"column:data_storage"`             // DataStorage 数据盘
	BandWith      int32             `json:"band_with" gorm:"band_width"`                         // BandWidth 公网 ip 带宽 单位 M
	Tags          map[string]string `json:"tags" gorm:"column:tags;serializer:json;type:json"`   // Tags 资源标签
	Extra         map[string]string `json:"extra" gorm:"column:extra;serializer:json;type:json"` // Extra 额外的通用属性
}

// Status 资源当前的状态
type Status struct {
	Phase          string   `json:"phase" gorm:"column:phase"`                                               // Phase 资源当前状态
	Describe       string   `json:"describe" gorm:"column:describe"`                                         // Describe 资源当前状态描述
	LockMode       string   `json:"lock_mode" gorm:"column:lock_mode"`                                       // LockMode 实例锁定模式 Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定
	LockReason     string   `json:"lock_reason" gorm:"column:lock_reason"`                                   // LockReason 锁定原因
	PublicAddress  []string `json:"public_address" gorm:"column:public_address;serializer:json;type:json"`   // PublicAddress 公网地址,或者域名 资源访问地址
	PrivateAddress []string `json:"private_address" gorm:"column:private_address;serializer:json;type:json"` // PrivateAddress 内网地址 或者域名
}

func (s *Status) GetFirstPrivateAddress() string {
	if len(s.PrivateAddress) > 0 {
		return s.PrivateAddress[0]
	}

	return ""
}

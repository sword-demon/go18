// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package policy

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/modules/iam/apps"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"time"
)

type Policy struct {
	apps.ResourceMeta
	CreatePolicyRequest

	Namespace *namespace.Namespace `json:"namespace,omitempty" gorm:"-"` // Namespace 关联空间
	User      *user.User           `json:"user,omitempty" gorm:"-"`      // User 关联用户
	Role      *role.Role           `json:"role,omitempty" gorm:"-"`      // Role 关联角色
}

func NewPolicy() *Policy {
	return &Policy{
		ResourceMeta: *apps.NewResourceMeta(),
	}
}

func (p *Policy) TableName() string {
	return "policy"
}

func (p *Policy) String() string {
	return pretty.ToJSON(p)
}

func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{
		Extras:   map[string]string{},
		Scope:    map[string]string{},
		Enabled:  true,
		ReadOnly: false,
	}
}

type CreatePolicyRequest struct {
	// 创建者
	CreateBy uint64 `json:"create_by" bson:"create_by" gorm:"column:create_by;type:uint" description:"创建者" optional:"true"`
	// 空间
	NamespaceId *uint64 `json:"namespace_id" bson:"namespace_id" gorm:"column:namespace_id;type:varchar(200);index" description:"策略生效的空间Id" optional:"true"`
	// 用户Id
	UserId uint64 `json:"user_id" bson:"user_id" gorm:"column:user_id;type:uint;not null;index" validate:"required" description:"被授权的用户"`
	// 角色Id
	RoleId uint64 `json:"role_id" bson:"role_id" gorm:"column:role_id;type:uint;not null;index" validate:"required" description:"被关联的角色"`
	// 访问范围, 需要提前定义scope, 比如环境
	Scope map[string]string `json:"scope" bson:"scope" gorm:"column:scope;serializer:json;type:json" description:"数据访问的范围" optional:"true"`
	// 策略过期时间
	ExpiredTime *time.Time `json:"expired_time" bson:"expired_time" gorm:"column:expired_time;type:timestamp;index" description:"策略过期时间" optional:"true"`
	// 只读策略, 不允许用户修改, 一般用于系统管理
	ReadOnly bool `json:"read_only" bson:"read_only" gorm:"column:read_only;type:tinyint(1)" description:"只读策略, 不允许用户修改, 一般用于系统管理" optional:"true"`
	// 该策略是否启用
	Enabled bool `json:"enabled" bson:"enabled" gorm:"column:enabled;type:tinyint(1)" description:"该策略是否启用" optional:"true"`
	// 策略标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index" description:"策略标签" optional:"true"`
	// 扩展信息
	Extras map[string]string `json:"extras" bson:"extras" gorm:"column:extras;serializer:json;type:json" description:"扩展信息" optional:"true"`
}

func (r *CreatePolicyRequest) Validate() error {
	return validator.Validate(r)
}

func (r *CreatePolicyRequest) SetNamespaceId(namespaceId uint64) *CreatePolicyRequest {
	r.NamespaceId = &namespaceId
	return r
}

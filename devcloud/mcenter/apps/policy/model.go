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
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

type Policy struct {
	// 基础数据
	apps.ResourceMeta
	// 策略定义
	CreatePolicyRequest

	// 关联空间
	Namespace *namespace.Namespace `json:"namespace,omitempty" gorm:"-"` // Namespace 关联空间
	// 关联用户
	User *user.User `json:"user,omitempty" gorm:"-"` // User 关联用户
	// 关联角色
	Role *role.Role `json:"role,omitempty" gorm:"-"` // Role 关联角色
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
		ResourceScope: ResourceScope{
			Scope: map[string][]string{},
		},
	}
}

type CreatePolicyRequest struct {
	ResourceScope
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

func (r *ResourceScope) GetNamespaceId() uint64 {
	if r.NamespaceId == nil {
		return 0
	}
	return *r.NamespaceId
}

func (r *ResourceScope) BuildMySQLPrefixBlob() {
	for k := range r.Scope {
		for i := range r.Scope[k] {
			if !strings.HasSuffix(r.Scope[k][i], "%") {
				r.Scope[k][i] += "%"
			}
		}
	}
}

// ResourceScope 资源需要组合 ResourceLabel 使用
type ResourceScope struct {
	NamespaceId *uint64             `json:"namespace_id" bson:"namespace_id" gorm:"column:namespace_id;type:varchar(200);index" description:"策略生效的空间 id" optional:"true"`
	Scope       map[string][]string `json:"scope" bson:"scope" gorm:"column:scope;serializer:json;type:json" description:"数据访问的范围" optional:"true"`
}

func (r *ResourceScope) SetScope(key string, value []string) {
	if r.Scope == nil {
		r.Scope = map[string][]string{}
	}
	r.Scope[key] = value
}

func (r *ResourceScope) GormResourceFilter(query *gorm.DB) *gorm.DB {
	if r.NamespaceId != nil {
		query = query.Where("namespace_id = ?", r.NamespaceId)
	}

	for key, values := range r.Scope {
		if len(values) == 0 {
			continue
		}

		// 构建 标签不存在 条件
		notHasKey := clause.Not(datatypes.JSONQuery("label").HasKey(key))

		// 构建 标签值匹配 条件
		var valueMatches []clause.Expression
		for _, val := range values {
			valueMatches = append(valueMatches,
				datatypes.JSONQuery("label").Equals(val, key))
		}

		// 组合条件:标签不存在 or 标签匹配值
		query = query.Where(clause.Or(
			notHasKey,
			clause.Or(valueMatches...)),
		)
	}

	return query
}

type ResourceLabel struct {
	NamespaceId *uint64           `json:"namespace_id" bson:"namespace_id" gorm:"column:namespace_id;type:varchar(200);index" description:"策略生效的空间 id" optional:"true"`
	Label       map[string]string `json:"label" bson:"label" gorm:"column:label;serializer:json;type:json" description:"资源标签, 用于标识资源的属性" optional:"true"`
}

func (l *ResourceLabel) SetLabel(key, value string) {
	if l.Label == nil {
		l.Label = map[string]string{}
	}
	l.Label[key] = value
}

func (l *ResourceLabel) SetNamespaceId(v uint64) {
	l.NamespaceId = &v
}

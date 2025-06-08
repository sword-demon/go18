// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package role

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/modules/iam/apps"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
)

type Role struct {
	// 基础数据
	apps.ResourceMeta
	// 角色创建信息
	CreateRoleRequest
	// 菜单权限
	MenuPermissions []*ViewPermission `json:"menu_permissions,omitempty" gorm:"-" description:"角色关联的菜单权限"`
	// API 权限
	ApiPermissions []*ApiPermission `json:"api_permissions,omitempty" gorm:"-" description:"角色关联的 API 权限"`
}

func NewRole() *Role {
	return &Role{
		ResourceMeta:    *apps.NewResourceMeta(),
		MenuPermissions: []*ViewPermission{},
		ApiPermissions:  []*ApiPermission{},
	}
}

func (r *Role) TableName() string {
	return "roles"
}
func (r *Role) String() string {
	return pretty.ToJSON(r)
}

// CheckPerm 该角色是否允许该API访问
func (r *Role) CheckPerm(re *endpoint.RouteEntry) error {
	return nil
}

func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Extras:  map[string]string{},
		Enabled: true,
	}
}

type CreateRoleRequest struct {
	// 创建者ID
	CreateBy uint64 `json:"create_by" gorm:"column:create_by" description:"创建者ID" optional:"true"`
	// 角色名称
	Name string `json:"name" gorm:"column:name;type:varchar(100);index" bson:"name" description:"角色名称"`
	// 角色描述
	Description string `json:"description" gorm:"column:description;type:text" bson:"description" description:"角色描述"`
	// 是否启用
	Enabled bool `json:"enabled" bson:"enabled" gorm:"column:enabled;type:tinyint(1)" description:"是否启用" optional:"true"`
	// 标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index" description:"标签" optional:"true"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" description:"其他扩展信息" optional:"true"`
}

func (r *CreateRoleRequest) Validate() error {
	return validator.Validate(r)
}

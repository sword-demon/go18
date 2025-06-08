// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package role

import (
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/modules/iam/apps"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
)

type ApiPermission struct {
	apps.ResourceMeta
	RoleId uint64 `json:"role_id" gorm:"column:role_id;index" description:"角色 id"`
	// Api 权限定义
	ApiPermissionSpec
}

func NewApiPermission(roleId uint64, apiPermissionSpec *ApiPermissionSpec) *ApiPermission {
	return &ApiPermission{RoleId: roleId, ApiPermissionSpec: *apiPermissionSpec, ResourceMeta: *apps.NewResourceMeta()}
}

type ApiPermissionSpec struct {
	// 创建者ID
	CreateBy uint64 `json:"create_by" gorm:"column:create_by" description:"创建者ID" optional:"true"`
	// 角色描述
	Description string `json:"description" gorm:"column:description;type:text" bson:"description" description:"角色描述"`
	// 权限匹配方式
	MatchBy MatchBy `json:"match_by" gorm:"column:match_by;type:tinyint(1);index" bson:"match_by" description:"权限匹配方式"`
	// MATCH_BY_ID 时指定的 Endpoint Id
	EndpointId *uint64 `json:"endpoint_id" gorm:"column:endpoint_id;type:uint;index"`
	// 操作标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index"`
	// 服务
	Service string `json:"service" gorm:"column:service;type:varchar(100);index" bson:"service" description:"服务名称"`
	// 资源列表
	Resource string `json:"resource" gorm:"column:resource;type:varchar(100);index" bson:"resource" description:"资源名称"`
	// 资源操作
	Action string `json:"action" bson:"action" gorm:"column:action;type:varchar(100);index"`
	// 读或者读写
	AccessMode endpoint.AccessMode `json:"access_mode" bson:"access_mode" gorm:"column:access_mode;type:tinyint(1);index"`

	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" description:"其他扩展信息" optional:"true"`
}

func (r *ApiPermission) TableName() string {
	return "api_permissions"
}

func (r *ApiPermission) String() string {
	return pretty.ToJSON(r)
}

func NewResourceActionApiPermissionSpec(service, resource, action string) *ApiPermissionSpec {
	return &ApiPermissionSpec{
		Extras:   map[string]string{},
		MatchBy:  MatchByResourceAction,
		Service:  service,
		Resource: resource,
		Action:   action,
	}
}

func (a *ApiPermissionSpec) GetEndpointId() uint64 {
	if a.EndpointId == nil {
		return 0
	}

	return *a.EndpointId
}

// IsMatch 判断是否有当前 API 的访问权限
func (a *ApiPermissionSpec) IsMatch(target *endpoint.Endpoint) bool {
	switch a.MatchBy {
	case MatchById:
		if a.EndpointId == nil {
			return false
		}
		if *a.EndpointId == target.Id {
			return true
		}
	case MatchByResourceAccessMode:
		if a.AccessMode == target.AccessMode {
			return true
		}
	case MatchByResourceAction:
		if a.Service != "*" && a.Service != target.Service {
			return false
		}
		if a.Resource != "*" && a.Resource != target.Resource {
			return false
		}
		if a.Action != "*" && a.Action != target.Action {
			return false
		}
		return true
	default:
		return false
	}

	return false
}

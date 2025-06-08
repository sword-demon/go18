package role

import "github.com/infraboard/modules/iam/apps"

type ViewPermission struct {
	apps.ResourceMeta
	RoleId uint64 `json:"role_id" gorm:"column:role_id;index" description:"Role_id"`
	// menu 权限定义
	ViewPermissionSpec
}

func NewViewPermission(roleId uint64, spec *ViewPermissionSpec) *ViewPermission {
	return &ViewPermission{RoleId: roleId, ViewPermissionSpec: *spec, ResourceMeta: *apps.NewResourceMeta()}
}

func (r *ViewPermission) TableName() string {
	return "view_permissions"
}

func NewViewPermissionSpec() *ViewPermissionSpec {
	return &ViewPermissionSpec{
		Extras: map[string]string{},
	}
}

type ViewPermissionSpec struct {
	// CreateBy 创建者 id
	CreateBy uint64 `json:"create_by" gorm:"column:create_by" description:"创建者 id" optional:"true"`
	// Description 描述
	Description string `json:"description" gorm:"column:description;type:text" bson:"description" description:"角色描述"`
	// 页面路径
	PagePath string `json:"path_path" gorm:"column:path_path;type:varchar(200);index" bson:"path_path" description:"页面路径(可以通配)"`
	// 组件名称
	Components []string `json:"components" gorm:"column:components;type:json;serializer:json" bson:"components" description:"页面组件(可以通配)"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" description:"其他扩展信息" optional:"true"`
}

package role

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
	"github.com/infraboard/modules/iam/apps/view"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
)

const AppName = "role"

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	RoleService
	ApiPermissionService
	ViewPermissionService
}

type RoleService interface {
	// CreateRole 创建角色
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	// QueryRole 列表查询
	QueryRole(context.Context, *QueryRoleRequest) (*types.Set[*Role], error)
	// DescribeRole 详情查询
	DescribeRole(context.Context, *DescribeRoleRequest) (*Role, error)
	// UpdateRole 更新角色
	UpdateRole(context.Context, *UpdateRoleRequest) (*Role, error)
	// DeleteRole 删除角色
	DeleteRole(context.Context, *DeleteRoleRequest) (*Role, error)
}

type QueryRoleRequest struct {
	*request.PageRequest
	WithMenuPermission bool     `json:"with_menu_permission" form:"with_menu_permission"`
	WithApiPermission  bool     `json:"with_api_permission" form:"with_api_permission"`
	RoleIds            []uint64 `json:"role_ids" form:"role_ids"`
}

func NewQueryRoleRequest() *QueryRoleRequest {
	return &QueryRoleRequest{
		PageRequest: request.NewDefaultPageRequest(),
		RoleIds:     []uint64{},
	}
}

type DescribeRoleRequest struct {
	apps.GetRequest
}

func NewDescribeRoleRequest() *DescribeRoleRequest {
	return &DescribeRoleRequest{}
}

type UpdateRoleRequest struct {
	apps.GetRequest
	CreateRoleRequest
}

type DeleteRoleRequest struct {
	apps.GetRequest
}

func NewDeleteRoleRequest() *DeleteRoleRequest {
	return &DeleteRoleRequest{}
}

type ApiPermissionService interface {
	// QueryApiPermission 查询角色关联的权限条目
	QueryApiPermission(context.Context, *QueryApiPermissionRequest) ([]*ApiPermission, error)
	// AddApiPermission 添加角色关联API
	AddApiPermission(context.Context, *AddApiPermissionRequest) ([]*ApiPermission, error)
	// RemoveApiPermission 移除角色关联API
	RemoveApiPermission(context.Context, *RemoveApiPermissionRequest) ([]*ApiPermission, error)
	// QueryMatchedEndpoint 查询匹配到的Api接口列表
	QueryMatchedEndpoint(context.Context, *QueryMatchedEndpointRequest) (*types.Set[*endpoint.Endpoint], error)
}

type QueryApiPermissionRequest struct {
	RoleIds          []uint64 `json:"role_ids"`
	ApiPermissionIds []uint64 `json:"api_permission_ids"`
}

func NewQueryApiPermissionRequest() *QueryApiPermissionRequest {
	return &QueryApiPermissionRequest{
		RoleIds:          []uint64{},
		ApiPermissionIds: []uint64{},
	}
}

func (q *QueryApiPermissionRequest) AddRoleId(roleIds ...uint64) *QueryApiPermissionRequest {
	q.RoleIds = append(q.RoleIds, roleIds...)
	return q
}

func (q *QueryApiPermissionRequest) AddPermissionId(permissionIds ...uint64) *QueryApiPermissionRequest {
	q.ApiPermissionIds = append(q.ApiPermissionIds, permissionIds...)
	return q
}

type AddApiPermissionRequest struct {
	RoleId uint64               `json:"role_id"`
	Items  []*ApiPermissionSpec `json:"items"`
}

func NewAddApiPermissionRequest(roleId uint64) *AddApiPermissionRequest {
	return &AddApiPermissionRequest{RoleId: roleId}
}

func (a *AddApiPermissionRequest) Validate() error {
	return validator.Validate(a)
}

func (a *AddApiPermissionRequest) Add(specs ...*ApiPermissionSpec) *AddApiPermissionRequest {
	a.Items = append(a.Items, specs...)
	return a
}

type RemoveApiPermissionRequest struct {
	RoleId           uint64   `json:"role_id"`
	ApiPermissionIds []uint64 `json:"api_permission_ids"`
}

func NewRemoveApiPermissionRequest(roleId uint64) *RemoveApiPermissionRequest {
	return &RemoveApiPermissionRequest{
		RoleId:           roleId,
		ApiPermissionIds: []uint64{},
	}
}

func (r *RemoveApiPermissionRequest) Add(apiPermissionIds ...uint64) *RemoveApiPermissionRequest {
	r.ApiPermissionIds = append(r.ApiPermissionIds, apiPermissionIds...)
	return r
}

func (r *RemoveApiPermissionRequest) Validate() error {
	return validator.Validate(r)
}

type QueryMatchedEndpointRequest struct {
	RoleIds []uint64 `json:"role_ids" form:"role_ids"`
}

func NewQueryMatchedEndpointRequest() *QueryMatchedEndpointRequest {
	return &QueryMatchedEndpointRequest{RoleIds: []uint64{}}
}

// ViewPermissionService 角色菜单管理
type ViewPermissionService interface {
	// QueryViewPermission 查询角色关联的视图权限
	QueryViewPermission(context.Context, *QueryViewPermissionRequest) ([]*ViewPermission, error)
	// AddViewPermission 添加角色关联菜单
	AddViewPermission(context.Context, *AddViewPermissionRequest) ([]*ViewPermission, error)
	// RemoveViewPermission 移除角色关联菜单
	RemoveViewPermission(context.Context, *RemoveViewPermissionRequest) ([]*ViewPermission, error)
	// QueryMatchedPage 查询能匹配到视图菜单
	QueryMatchedPage(context.Context, *QueryMatchedPageRequest) (*types.Set[*view.Menu], error)
}

type QueryViewPermissionRequest struct {
	RoleIds           []uint64 `json:"role_ids"`
	ViewPermissionIds []uint64 `json:"view_permission_ids"`
}

func NewQueryViewPermissionRequest() *QueryViewPermissionRequest {
	return &QueryViewPermissionRequest{
		RoleIds:           []uint64{},
		ViewPermissionIds: []uint64{},
	}
}

func (q *QueryViewPermissionRequest) AddRoleId(roleIds ...uint64) *QueryViewPermissionRequest {
	q.RoleIds = append(q.RoleIds, roleIds...)
	return q
}

func (q *QueryViewPermissionRequest) AddPermissionId(permissionIds ...uint64) *QueryViewPermissionRequest {
	q.ViewPermissionIds = append(q.ViewPermissionIds, permissionIds...)
	return q
}

type AddViewPermissionRequest struct {
	RoleId uint64                `json:"role_id"`
	Items  []*ViewPermissionSpec `json:"items"`
}

func NewAddViewPermissionRequest() *AddViewPermissionRequest {
	return &AddViewPermissionRequest{Items: []*ViewPermissionSpec{}}
}

func (a *AddViewPermissionRequest) Validate() error {
	return validator.Validate(a)
}
func (a *AddViewPermissionRequest) Add(specs ...*ViewPermissionSpec) *AddViewPermissionRequest {
	a.Items = append(a.Items, specs...)
	return a
}

type RemoveViewPermissionRequest struct {
	RoleId            uint64   `json:"role_id"`
	ViewPermissionIds []uint64 `json:"view_permission_ids"`
}

func NewRemoveViewPermissionRequest() *RemoveViewPermissionRequest {
	return &RemoveViewPermissionRequest{ViewPermissionIds: []uint64{}}
}

func (r *RemoveViewPermissionRequest) Validate() error {
	return validator.Validate(r)
}

type QueryMatchedPageRequest struct {
	apps.GetRequest
}

func NewQueryMatchedPageRequest() *QueryMatchedPageRequest {
	return &QueryMatchedPageRequest{}
}

package policy

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
	"github.com/infraboard/modules/iam/apps/view"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
)

const (
	AppName = "policy"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	PolicyService
	PermissionService
}

type PolicyService interface {
	// CreatePolicy 创建策略
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	// QueryPolicy 查询策略列表
	QueryPolicy(context.Context, *QueryPolicyRequest) (*types.Set[*Policy], error)
	// DescribePolicy 查询详情
	DescribePolicy(context.Context, *DescribePolicyRequest) (*Policy, error)
	// UpdatePolicy 更新策略
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error)
	// DeletePolicy 删除策略
	DeletePolicy(context.Context, *DeletePolicyRequest) (*Policy, error)
}

type QueryPolicyRequest struct {
	*request.PageRequest
	// 忽略分页
	SkipPage bool `json:"skip_page"`
	// 关联用户 id
	UserId *uint64 `json:"user_id"`
	// 关联空间 id
	NamespaceId *uint64 `json:"namespace_id"`
	// 是否过期
	Expired *bool `json:"expired"`
	// 有没有启用
	Enabled       *bool `json:"enabled"`
	WithNamespace bool  `json:"with_namespace"`
	// 关联查询出用户对象
	WithUser bool `json:"with_user"`
	// 关联查询角色对象
	WithRole bool `json:"with_role"`
}

func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{PageRequest: request.NewDefaultPageRequest()}
}

func (r *QueryPolicyRequest) SetNamespaceId(nsId uint64) *QueryPolicyRequest {
	r.NamespaceId = &nsId
	return r
}

func (r *QueryPolicyRequest) SetUserId(uid uint64) *QueryPolicyRequest {
	r.UserId = &uid
	return r
}

func (r *QueryPolicyRequest) SetExpired(v bool) *QueryPolicyRequest {
	r.Expired = &v
	return r
}

func (r *QueryPolicyRequest) SetEnabled(v bool) *QueryPolicyRequest {
	r.Enabled = &v
	return r
}

func (r *QueryPolicyRequest) SetSkipPage(v bool) *QueryPolicyRequest {
	r.SkipPage = v
	return r
}

func (r *QueryPolicyRequest) SetWithRole(v bool) *QueryPolicyRequest {
	r.WithRole = v
	return r
}
func (r *QueryPolicyRequest) SetWithUsers(v bool) *QueryPolicyRequest {
	r.WithUser = v
	return r
}
func (r *QueryPolicyRequest) SetWithUser(v bool) *QueryPolicyRequest {
	r.WithNamespace = v
	return r
}

type DescribePolicyRequest struct {
	apps.GetRequest
}

func NewDescribePolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{}
}

type UpdatePolicyRequest struct {
	apps.GetRequest
	CreatePolicyRequest
}

type DeletePolicyRequest struct {
	apps.GetRequest
}

func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{}
}

type PermissionService interface {
	// QueryNamespace 查询用户可以访问的空间
	QueryNamespace(context.Context, *QueryNamespaceRequest) (*types.Set[*namespace.Namespace], error)
	// QueryMenu 查询用户可以访问的菜单
	QueryMenu(context.Context, *QueryMenuRequest) (*types.Set[*view.Menu], error)
	// QueryEndpoint 查询用户可以访问的Api接口
	QueryEndpoint(context.Context, *QueryEndpointRequest) (*types.Set[*endpoint.Endpoint], error)
	// ValidatePagePermission 校验页面权限
	ValidatePagePermission(context.Context, *ValidatePagePermissionRequest) (*ValidatePagePermissionResponse, error)
	// ValidateEndpointPermission 校验接口权限
	ValidateEndpointPermission(context.Context, *ValidateEndpointPermissionRequest) (*ValidateEndpointPermissionResponse, error)
}

type ValidatePagePermissionRequest struct {
	UserId      uint64 `json:"user_id" form:"user_id"`
	NamespaceId uint64 `json:"namespace_id" form:"namespace_id"`
	Path        string `json:"path" form:"path"`
}

func NewValidatePagePermissionResponse(req ValidatePagePermissionRequest) *ValidatePagePermissionResponse {
	return &ValidatePagePermissionResponse{
		ValidatePagePermissionRequest: req,
	}
}

type ValidatePagePermissionResponse struct {
	ValidatePagePermissionRequest
	HasPermission bool       `json:"has_permission"`
	Page          *view.Page `json:"page"`
}

type ValidateEndpointPermissionRequest struct {
	UserId      uint64 `json:"user_id" form:"user_id"`
	NamespaceId uint64 `json:"namespace_id" form:"namespace_id"`
	Service     string `json:"service" form:"service"`
	Path        string `json:"path" form:"path"`
	Method      string `json:"method" form:"method"`
}

func NewValidateEndpointPermissionRequest() *ValidateEndpointPermissionRequest {
	return &ValidateEndpointPermissionRequest{}
}

type ValidateEndpointPermissionResponse struct {
	ValidateEndpointPermissionRequest
	HasPermission bool               `json:"has_permission"`
	Endpoint      *endpoint.Endpoint `json:"endpoint"`
}

func NewValidateEndpointPermissionResponse(req ValidateEndpointPermissionRequest) *ValidateEndpointPermissionResponse {
	return &ValidateEndpointPermissionResponse{ValidateEndpointPermissionRequest: req}
}

type QueryNamespaceRequest struct {
	UserId      uint64 `json:"user_id"`
	NamespaceId uint64 `json:"namespace_id"`
}

func NewQueryNamespaceRequest() *QueryNamespaceRequest {
	return &QueryNamespaceRequest{}
}

func (r *QueryNamespaceRequest) SetUserId(v uint64) *QueryNamespaceRequest {
	r.UserId = v
	return r
}

func (r *QueryNamespaceRequest) SetNamespaceId(v uint64) *QueryNamespaceRequest {
	r.NamespaceId = v
	return r
}

func NewQueryMenuRequest() *QueryMenuRequest {
	return &QueryMenuRequest{}
}

type QueryMenuRequest struct {
	UserId      uint64 `json:"user_id"`
	NamespaceId uint64 `json:"namespace_id"`
}

type QueryEndpointRequest struct {
	UserId      uint64 `json:"user_id"`
	NamespaceId uint64 `json:"namespace_id"`
}

func NewQueryEndpointRequest() *QueryEndpointRequest {
	return &QueryEndpointRequest{}
}

func (r *QueryEndpointRequest) SetUserId(v uint64) *QueryEndpointRequest {
	r.UserId = v
	return r
}

func (r *QueryEndpointRequest) SetNamespaceId(v uint64) *QueryEndpointRequest {
	r.NamespaceId = v
	return r
}

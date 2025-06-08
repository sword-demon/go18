package endpoint

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/google/uuid"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
)

type Endpoint struct {
	apps.ResourceMeta
	RouteEntry `bson:",inline" validate:"required"`
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		ResourceMeta: *apps.NewResourceMeta(),
	}
}

func IsEndpointExist(set *types.Set[*Endpoint], target *Endpoint) bool {
	for _, item := range set.Items {
		if item.Id == target.Id {
			return true
		}
	}

	return false
}

func (e *Endpoint) TableName() string {
	return "endpoints"
}

func (e *Endpoint) String() string {
	return pretty.ToJSON(e)
}

func (e *Endpoint) IsMatched(service, method, path string) bool {
	if e.Service != service {
		return false
	}
	if e.Method != method {
		return false
	}
	if e.Path != path {
		return false
	}
	return true
}

func (e *Endpoint) SetRouteEntry(v RouteEntry) *Endpoint {
	e.RouteEntry = v
	return e
}

// RouteEntry 路由条目
type RouteEntry struct {
	// 该功能属于那个服务
	UUID string `json:"uuid" bson:"uuid" gorm:"column:uuid;type:varchar(100);uniqueIndex" optional:"true" description:"路由UUID"`
	// 该功能属于那个服务
	Service string `json:"service" bson:"service" validate:"required,lte=64" gorm:"column:service;type:varchar(100);index" description:"服务名称"`
	// 服务那个版本的功能
	Version string `json:"version" bson:"version" validate:"required,lte=64" gorm:"column:version;type:varchar(100)" optional:"true" description:"版本版本"`
	// 资源名称
	Resource string `json:"resource" bson:"resource" gorm:"column:resource;type:varchar(100);index" description:"资源名称"`
	// 资源操作
	Action string `json:"action" bson:"action" gorm:"column:action;type:varchar(100);index" description:"资源操作"`
	// 读或者写
	AccessMode AccessMode `json:"access_mode" bson:"access_mode" gorm:"column:access_mode;type:tinyint(1);index" optional:"true" description:"读写权限"`
	// 操作标签
	ActionLabel string `json:"action_label" gorm:"column:action_label;type:varchar(200);index" optional:"true" description:"资源标签"`
	// 函数名称
	FunctionName string `json:"function_name" bson:"function_name" gorm:"column:function_name;type:varchar(100)"  optional:"true" description:"函数名称"`
	// HTTP path 用于自动生成http api
	Path string `json:"path" bson:"path" gorm:"column:path;type:varchar(200);index" description:"接口的路径"`
	// HTTP method 用于自动生成http api
	Method string `json:"method" bson:"method" gorm:"column:method;type:varchar(100);index" description:"接口的方法"`
	// 接口说明
	Description string `json:"description" bson:"description" gorm:"column:description;type:text" optional:"true" description:"接口说明"`
	// 是否校验用户身份 (acccess_token 校验)
	RequiredAuth bool `json:"required_auth" bson:"required_auth" gorm:"column:required_auth;type:tinyint(1)" optional:"true" description:"是否校验用户身份 (acccess_token 校验)"`
	// 验证码校验(开启双因子认证需要) (code 校验)
	RequiredCode bool `json:"required_code" bson:"required_code" gorm:"column:required_code;type:tinyint(1)" optional:"true" description:"验证码校验(开启双因子认证需要) (code 校验)"`
	// 开启鉴权
	RequiredPerm bool `json:"required_perm" bson:"required_perm" gorm:"column:required_perm;type:tinyint(1)" optional:"true" description:"开启鉴权"`
	// ACL模式下, 允许的通过的身份标识符, 比如角色, 用户类型之类
	RequiredRole []string `json:"required_role" bson:"required_role" gorm:"column:required_role;serializer:json;type:json" optional:"true" description:"ACL模式下, 允许的通过的身份标识符, 比如角色, 用户类型之类"`
	// 是否开启操作审计, 开启后这次操作将被记录
	RequiredAudit bool `json:"required_audit" bson:"required_audit" gorm:"column:required_audit;type:tinyint(1)" optional:"true" description:"是否开启操作审计, 开启后这次操作将被记录"`
	// 名称空间不能为空
	RequiredNamespace bool `json:"required_namespace" bson:"required_namespace" gorm:"column:required_namespace;type:tinyint(1)" optional:"true" description:"名称空间不能为空"`
	// 扩展信息
	Extras map[string]string `json:"extras" bson:"extras" gorm:"column:extras;serializer:json;type:json" optional:"true" description:"扩展信息"`
}

func NewRouteEntry() *RouteEntry {
	return &RouteEntry{
		RequiredRole: []string{},
		Extras:       make(map[string]string),
	}
}

func (r *RouteEntry) BuildUUID() *RouteEntry {
	r.UUID = uuid.NewSHA1(uuid.Nil, fmt.Appendf(nil, "%s-%s-%s", r.Service, r.Method, r.Path)).String()
	return r
}

func GetRouteMeta[T any](m map[string]any, key string) T {
	if v, ok := m[key]; ok {
		return v.(T)
	}

	var t T
	return t
}

func (r *RouteEntry) LoadMeta(meta map[string]any) {
	r.Service = application.Get().AppName
	r.Resource = GetRouteMeta[string](meta, MetaResourceKey)
	r.Action = GetRouteMeta[string](meta, MetaActionKey)
	r.RequiredAuth = GetRouteMeta[bool](meta, MetaRequiredAuthKey)
	r.RequiredCode = GetRouteMeta[bool](meta, MetaRequiredCodeKey)
	r.RequiredPerm = GetRouteMeta[bool](meta, MetaRequiredPermKey)
	r.RequiredRole = GetRouteMeta[[]string](meta, MetaRequiredRoleKey)
	r.RequiredAudit = GetRouteMeta[bool](meta, MetaRequiredAuditKey)
	r.RequiredNamespace = GetRouteMeta[bool](meta, MetaRequiredNamespaceKey)
}

func (r *RouteEntry) HasRequiredRole() bool {
	return len(r.RequiredRole) > 0
}

// UniquePath todo
func (r *RouteEntry) UniquePath() string {
	return fmt.Sprintf("%s.%s", r.Method, r.Path)
}

func (r *RouteEntry) IsRequireRole(target string) bool {
	for i := range r.RequiredRole {
		if r.RequiredRole[i] == "*" {
			return true
		}

		if r.RequiredRole[i] == target {
			return true
		}
	}

	return false
}

func (r *RouteEntry) SetRequiredAuth(v bool) *RouteEntry {
	r.RequiredAuth = v
	return r
}

func (r *RouteEntry) AddRequiredRole(roles ...string) *RouteEntry {
	r.RequiredRole = append(r.RequiredRole, roles...)
	return r
}

func (r *RouteEntry) SetRequiredPerm(v bool) *RouteEntry {
	r.RequiredPerm = v
	return r
}

func (r *RouteEntry) SetLabel(value string) *RouteEntry {
	r.ActionLabel = value
	return r
}

func (r *RouteEntry) SetExtensionFromMap(m map[string]string) *RouteEntry {
	if r.Extras == nil {
		r.Extras = map[string]string{}
	}

	for k, v := range m {
		r.Extras[k] = v
	}
	return r
}

func (r *RouteEntry) SetRequiredCode(v bool) *RouteEntry {
	r.RequiredCode = v
	return r
}

func NewEntryFromRestRequest(req *restful.Request) *RouteEntry {
	entry := NewRouteEntry()

	// 请求拦截
	route := req.SelectedRoute()
	if route == nil {
		return nil
	}

	entry.FunctionName = route.Operation()
	entry.Method = route.Method()
	entry.LoadMeta(route.Metadata())
	entry.Path = route.Path()
	return entry
}

func NewEntryFromRestRouteReader(route restful.RouteReader) *RouteEntry {
	entry := NewRouteEntry()
	entry.FunctionName = route.Operation()
	entry.Method = route.Method()
	entry.LoadMeta(route.Metadata())
	entry.Path = route.Path()
	return entry
}

func NewEntryFromRestRoute(route restful.Route) *RouteEntry {
	entry := NewRouteEntry()
	entry.FunctionName = route.Operation
	entry.Method = route.Method
	entry.LoadMeta(route.Metadata)
	entry.Path = route.Path
	return entry
}

// NewEntryFromRestfulContainer 获取 container 里的所有路由条目
func NewEntryFromRestfulContainer(c *restful.Container) (entries []*RouteEntry) {
	// 获取当前 container 里的所有 webService
	wss := c.RegisteredWebServices()
	for i := range wss {
		// 获取当前 webService 的所有路由
		routes := wss[i].Routes()
		for _, route := range routes {
			// 将当前路由转换为 RouteEntry 实例
			es := NewEntryFromRestRoute(route)
			if es != nil {
				// 将非 nil 的 RouteEntry 实例添加到结果切片中
				entries = append(entries, es)
			}
		}
	}
	return entries
}

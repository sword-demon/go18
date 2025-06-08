package endpoint

import (
	"context"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
	"slices"
)

const (
	AppName = "endpoint"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	// RegistryEndpoint 注册 api 接口
	RegistryEndpoint(context.Context, *RegistryEndpointRequest) (*types.Set[*Endpoint], error)
	// QueryEndpoint 查询 API 接口列表
	QueryEndpoint(context.Context, *QueryEndpointRequest) (*types.Set[*Endpoint], error)
	// DescribeEndpoint 查询 API 接口详情
	DescribeEndpoint(context.Context, *DescribeEndpointRequest) (*Endpoint, error)
}

type RegistryEndpointRequest struct {
	Items []*RouteEntry `json:"items"`
}

func NewRegistryEndpointRequest() *RegistryEndpointRequest {
	return &RegistryEndpointRequest{
		Items: []*RouteEntry{},
	}
}

type QueryEndpointRequest struct {
	Services []string `form:"services" json:"services"`
}

func NewQueryEndpointRequest() *QueryEndpointRequest {
	return &QueryEndpointRequest{}
}

func (r *QueryEndpointRequest) IsMatchAllService() bool {
	return slices.Contains(r.Services, "*")
}

func (r *QueryEndpointRequest) WithService(services ...string) *QueryEndpointRequest {
	for _, service := range services {
		if !slices.Contains(r.Services, service) {
			r.Services = append(r.Services, services...)
		}
	}
	return r
}

type DescribeEndpointRequest struct {
	apps.GetRequest
}

func (r *RegistryEndpointRequest) AddItem(items ...*RouteEntry) *RegistryEndpointRequest {
	r.Items = append(r.Items, items...)
	return r
}

func (r *RegistryEndpointRequest) Validate() error {
	return validator.Validate(r)
}

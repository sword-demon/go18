package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
)

func init() {
	ioc.Controller().Registry(&EndpointServiceImpl{})
}

var _ endpoint.Service = (*EndpointServiceImpl)(nil)

type EndpointServiceImpl struct {
	ioc.ObjectImpl
}

func (s *EndpointServiceImpl) Name() string {
	return endpoint.AppName
}

func (s *EndpointServiceImpl) Init() error {
	// 自动创建表
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&endpoint.Endpoint{})
		if err != nil {
			return err
		}
	}

	return nil
}

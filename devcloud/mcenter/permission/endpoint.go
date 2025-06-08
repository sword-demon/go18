package permission

import (
	"context"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
)

func init() {
	ioc.Api().Registry(&ApiRegister{})
}

type ApiRegister struct {
	ioc.ObjectImpl

	log *zerolog.Logger
}

func (a *ApiRegister) Name() string {
	return "api_register"
}

// Priority 这个 Init 一定要放在所有路由都添加完成之后运行
func (a *ApiRegister) Priority() int {
	return -100
}

func (a *ApiRegister) Init() error {
	a.log = log.Sub(a.Name())
	// 注册认证中间件
	entries := endpoint.NewEntryFromRestfulContainer(gorestful.RootRouter())
	req := endpoint.NewRegistryEndpointRequest()
	req.AddItem(entries...)
	set, err := endpoint.GetService().RegistryEndpoint(context.Background(), req)
	if err != nil {
		a.log.Error().Err(err).Msg("registry endpoint")
		return err
	}
	a.log.Info().Msgf("registry endpoints: %s", set.Items)
	return nil
}

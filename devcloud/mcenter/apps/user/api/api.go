package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
)

func init() {
	// 注册用户模块服务
	ioc.Api().Registry(&UserRestfulApiHandler{})
}

type UserRestfulApiHandler struct {
	ioc.ObjectImpl

	svc user.Service // 依赖 user 服务实现
}

func (h *UserRestfulApiHandler) Name() string {
	return user.AppName
}

func (h *UserRestfulApiHandler) Init() error {
	h.svc = user.GetService() // 获取服务

	tags := []string{"用户接口"} // 文档的 tag
	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("").To(h.CreateUser).
		Doc("创建用户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(user.User{}).
		Returns(200, "OK", user.User{}))

	return nil
}

package api

import (
	_ "embed"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
)

func init() {
	ioc.Api().Registry(&TokenRestfulApiHandler{})
}

type TokenRestfulApiHandler struct {
	ioc.ObjectImpl

	svc token.Service // 依赖 token 服务实现
}

func (h *TokenRestfulApiHandler) Name() string {
	return token.AppName
}

//go:embed docs/login.md
var loginApiDocNotes string

func (h *TokenRestfulApiHandler) Init() error {
	h.svc = token.GetService() // 获取服务

	tags := []string{"用户登录"} // 文档的 tag
	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("").To(h.Login).
		Doc("颁发令牌(登录)").
		Notes(loginApiDocNotes).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.POST("/validate").To(h.ValidateToken).
		Doc("校验令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.DELETE("").To(h.Logout).
		Doc("撤销令牌(退出)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}).
		Returns(404, "Not Found", nil))

	return nil
}

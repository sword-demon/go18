package private_token

import (
	"context"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
)

func init() {
	ioc.Config().Registry(&PrivateTokenIssuer{})
}

type PrivateTokenIssuer struct {
	ioc.ObjectImpl

	user  user.Service  // 依赖用户模块
	token token.Service // 依赖 token 模块
}

func (p *PrivateTokenIssuer) IssueToken(ctx context.Context, parameter token.IssueParameter) (*token.Token, error) {
	// 校验 token 的合法性
	oldTk, err := p.token.ValidateToken(ctx, token.NewValidateTokenRequest(parameter.AccessToken()))
	if err != nil {
		return nil, err
	}

	// 查询用户
	userReq := user.NewDescribeUserRequestById(oldTk.UserIdString())
	userM, err := p.user.DescribeUser(ctx, userReq)
	if err != nil {
		if exception.IsNotFoundError(err) {
			return nil, exception.NewUnauthorized("%s", err)
		}
		return nil, err
	}

	if !userM.EnabledApi {
		return nil, exception.NewPermissionDeny("未开启接口登录")
	}

	// 颁发 token
	tk := token.NewToken()
	tk.UserId = userM.Id
	tk.UserName = userM.UserName
	tk.IsAdmin = userM.IsAdmin

	expiredTTL := parameter.ExpireTTL()
	if expiredTTL > 0 {
		tk.SetExpiredAtByDuration(expiredTTL, 4)
	}

	return tk, nil
}

func (p *PrivateTokenIssuer) Name() string {
	return "private_token"
}

func (p *PrivateTokenIssuer) Init() error {
	p.user = user.GetService()
	p.token = token.GetService()

	token.RegistryIssuer(token.IssuerPrivateToken, p)
	return nil
}

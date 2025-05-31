package password

import (
	"context"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"time"
)

func init() {
	ioc.Config().Registry(&PwdTokenIssuer{
		ExpiredTTLSecond: 1 * 60 * 60,
	})
}

type PwdTokenIssuer struct {
	ioc.ObjectImpl

	user user.Service

	// ExpiredTTLSecond password 颁发的 token 的过期时间由系统配置,不允许用户自己设置
	ExpiredTTLSecond int `json:"expired_ttl_second" toml:"expired_ttl_second" yaml:"expired_ttl_second" env:"EXPIRED_TTL_SECOND"`

	expiredDuration time.Duration
}

func (t *PwdTokenIssuer) IssueToken(ctx context.Context, parameter token.IssueParameter) (*token.Token, error) {
	// 查询用户
	userReq := user.NewDescribeUserRequestByUserName(parameter.Username())
	userM, err := t.user.DescribeUser(ctx, userReq)
	if err != nil {
		if exception.IsNotFoundError(err) {
			return nil, exception.NewUnauthorized("%s", err)
		}
		return nil, err
	}

	// 比对密码
	err = userM.CheckPassword(parameter.Password())
	if err != nil {
		return nil, err
	}

	// 颁发 token
	tk := token.NewToken()
	tk.UserId = userM.Id
	tk.UserName = userM.UserName
	tk.IsAdmin = userM.IsAdmin

	tk.SetExpiredAtByDuration(t.expiredDuration, 4)
	return tk, nil
}

func (t *PwdTokenIssuer) Name() string {
	return "password_token_issuer"
}

func (t *PwdTokenIssuer) Init() error {
	t.user = user.GetService()
	t.expiredDuration = time.Duration(t.ExpiredTTLSecond) * time.Second

	token.RegistryIssuer(token.IssuerPassword, t)
	return nil
}

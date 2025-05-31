package token

import "context"

type Service interface {
	// IssueToken 颁发令牌: Login
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// RevokeToken 撤销令牌: Logout
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)
	// ValidateToken 验证令牌: 检查令牌的合法性,是否伪造
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

// IssueTokenRequest 用户的身份的凭证,用于换取token
type IssueTokenRequest struct {
	// Source 端类型
	Source SOURCE `json:"source"`
	// Issuer 认证方式
	Issuer string `json:"issuer"`
	// Parameter 参数
	Parameter IssueParameter `json:"parameter"`
}

type RevokeTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewRevokeTokenRequest(at string, rt string) *RevokeTokenRequest {
	return &RevokeTokenRequest{AccessToken: at, RefreshToken: rt}
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{AccessToken: accessToken}
}

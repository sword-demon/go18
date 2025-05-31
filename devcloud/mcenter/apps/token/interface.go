// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package token

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
	"time"
)

const (
	AppName = "token"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	// IssueToken 颁发令牌: Login
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// RevokeToken 撤销令牌: Logout
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)
	// ValidateToken 验证令牌: 检查令牌的合法性,是否伪造
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
	// QueryToken 查询已经颁发出去的 token
	QueryToken(context.Context, *QueryTokenRequest) (*types.Set[*Token], error)
	DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error)
}

type DescribeTokenRequest struct {
	DescribeBy    DescribeBy `json:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}

func NewDescribeTokenRequest(accessToken string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		DescribeBy:    DescribeByAccessToken,
		DescribeValue: accessToken,
	}
}

type QueryTokenRequest struct {
	*request.PageRequest
	Active  *bool    `json:"active"`   // 当前可用的没过期的 token
	Source  *SOURCE  `json:"source"`   // 用户来源
	UserIds []uint64 `json:"user_ids"` // user_ids
}

func NewQueryTokenRequest() *QueryTokenRequest {
	return &QueryTokenRequest{
		PageRequest: request.NewDefaultPageRequest(),
		UserIds:     []uint64{},
	}
}

func (r *QueryTokenRequest) SetActive(v bool) *QueryTokenRequest {
	r.Active = &v
	return r
}

func (r *QueryTokenRequest) SetSource(v SOURCE) *QueryTokenRequest {
	r.Source = &v
	return r
}

func (r *QueryTokenRequest) AddUserId(userIds ...uint64) *QueryTokenRequest {
	r.UserIds = append(r.UserIds, userIds...)
	return r
}

// IssueTokenRequest 用户的身份的凭证,用于换取token
type IssueTokenRequest struct {
	// Source 端类型
	Source SOURCE `json:"source"`
	// Issuer 认证方式
	Issuer string `json:"issuer.go"`
	// Parameter 参数
	Parameter IssueParameter `json:"parameter"`
}

func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		Parameter: make(IssueParameter),
	}
}

type RevokeTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewRevokeTokenRequest(at string, rt string) *RevokeTokenRequest {
	return &RevokeTokenRequest{AccessToken: at, RefreshToken: rt}
}

func (i *IssueTokenRequest) IssueByPassword(username, password string) {
	i.Issuer = IssuerPassword
	i.Parameter.SetUsername(username)
	i.Parameter.SetPassword(password)
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{AccessToken: accessToken}
}

type IssueParameter map[string]any

func NewIssueParameter() IssueParameter {
	return make(IssueParameter)
}

func GetIssueParameterValue[T any](p IssueParameter, key string) T {
	v := p[key]
	if v != nil {
		if value, ok := v.(T); ok {
			return value
		}
	}
	var zero T
	return zero
}

func (p IssueParameter) Username() string {
	return GetIssueParameterValue[string](p, "username")
}

func (p IssueParameter) Password() string {
	return GetIssueParameterValue[string](p, "password")
}

func (p IssueParameter) SetUsername(v string) IssueParameter {
	p["username"] = v
	return p
}

func (p IssueParameter) SetPassword(v string) IssueParameter {
	p["password"] = v
	return p
}

/*
private token issuer.go parameter
*/

func (p IssueParameter) AccessToken() string {
	return GetIssueParameterValue[string](p, "access_token")
}

func (p IssueParameter) ExpireTTL() time.Duration {
	return time.Second * time.Duration(GetIssueParameterValue[int64](p, "expired_ttl"))
}

func (p IssueParameter) SetAccessToken(v string) IssueParameter {
	p["access_token"] = v
	return p
}

func (p IssueParameter) SetExpireTTL(v int64) IssueParameter {
	p["expired_ttl"] = v
	return p
}

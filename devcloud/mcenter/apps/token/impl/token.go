// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/desense"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"time"
)

func (i *TokenServiceImpl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 颁发令牌的多种方式
	// user and password
	// ldap
	// feishu
	// dingding
	// weichat work

	issuer := token.GetIssuer(req.Issuer)
	if issuer == nil {
		return nil, exception.NewBadRequest("issuer %s not support", req.Issuer)
	}

	tk, err := issuer.IssueToken(ctx, req.Parameter)
	if err != nil {
		return nil, err
	}
	// token 设置颁发器和来源
	tk.SetIssuer(req.Issuer).SetSource(req.Source)

	// 判断当前数据库有没有已经存在的 token
	activeTokenQueryReq := token.NewQueryTokenRequest().AddUserId(tk.UserId).SetSource(req.Source).SetActive(true)
	tks, err := i.QueryToken(ctx, activeTokenQueryReq)
	if err != nil {
		return nil, err
	}

	switch req.Source {
	// 每个端只能有一个活跃登录
	case token.SourceWeb, token.SourceIos, token.SourcePc, token.SourceAndroid:
		if tks.Len() > 0 {
			i.log.Debug().Msgf("use exist active token: %s", desense.Default().DeSense(tk.AccessToken, "4", "3"))
			return tks.Items[0], nil
		}
	case token.SourceApi:
		if tks.Len() > int(i.MaxActiveApiToken) {
			return nil, exception.NewBadRequest("max active api token overflow")
		}
	default:
		panic("unhandled default case")
	}

	if tk.NamespaceId == 0 {
		tk.NamespaceId = 1 // 给一个默认值 1
	}

	// 保存 token
	if err := datasource.DBFromCtx(ctx).Create(tk).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

// RevokeToken 撤销 token 登出
func (i *TokenServiceImpl) RevokeToken(ctx context.Context, req *token.RevokeTokenRequest) (*token.Token, error) {
	tk, err := i.DescribeToken(ctx, token.NewDescribeTokenRequest(req.AccessToken))
	if err != nil {
		return nil, err
	}
	if err := tk.CheckRefreshToken(req.RefreshToken); err != nil {
		return nil, err
	}

	tk.Lock(token.LockTypeRevoke, "user revoke token")
	err = datasource.DBFromCtx(ctx).Model(&token.Token{}).Where("access_token = ?", req.AccessToken).
		Where("refresh_token = ?", req.RefreshToken).
		Updates(tk.Status.ToMap()).Error
	if err != nil {
		return nil, err
	}

	return tk, nil
}

func (i *TokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	// 查询 token 是不是我们系统颁发的
	tk := token.NewToken()
	err := datasource.DBFromCtx(ctx).Where("access_token = ?", req.AccessToken).First(tk).Error
	if err != nil {
		return nil, err
	}

	// 判断是否过期
	if err := tk.IsAccessTokenExpired(); err != nil {
		// 判断刷新 token 是否过期
		if err := tk.IsRefreshTokenExpired(); err != nil {
			// 这个也过期了,就真的要退出了
			return nil, err
		}

		// 如果开启了自动刷新
		if i.AutoRefresh {
			tk.SetRefreshAt(time.Now())
			tk.SetExpiredAtByDuration(i.refreshDuration, 4)
			if err := datasource.DBFromCtx(ctx).Save(tk).Error; err != nil {
				i.log.Error().Msgf("auto refresh token error, %s", err.Error())
			}
		}
		return nil, err
	}

	return tk, nil
}

func (i *TokenServiceImpl) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*types.Set[*token.Token], error) {
	set := types.New[*token.Token]()
	query := datasource.DBFromCtx(ctx).Model(&token.Token{})

	if req.Active != nil {
		if *req.Active {
			query = query.Where("lock_at IS NULL AND refresh_token_expired_at > ?", time.Now())
		} else {
			query = query.Where("lock_at IS NOT NULL OR refresh_token_expired_at <= ?", time.Now())
		}
	}

	if req.Source != nil {
		query = query.Where("source = ?", *req.Source)
	}
	if len(req.UserIds) > 0 {
		query = query.Where("user_id IN ?", req.UserIds)
	}

	// 查询总量
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (i *TokenServiceImpl) DescribeToken(ctx context.Context, req *token.DescribeTokenRequest) (*token.Token, error) {
	query := datasource.DBFromCtx(ctx)
	switch req.DescribeBy {
	case token.DescribeByAccessToken:
		query = query.Where("access_token = ?", req.DescribeValue)
	default:
		return nil, exception.NewBadRequest("unsupported describe type %d", req.DescribeBy)
	}

	tk := token.NewToken()
	if err := query.First(tk).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

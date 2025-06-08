// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package token

import "github.com/infraboard/mcube/v2/exception"

const (
	AccessTokenHeaderName         = "Authorization"
	AccessTokenCookieName         = "access_token"
	AccessTokenResponseHeaderName = "X-OAUTH-TOKEN"
	RefreshTokenHeaderName        = "X-REFRESH-TOKEN"
)

// 颁发器的类型
const (
	IssuerLDAP         = "ldap"
	IssuerFEISHU       = "feishu"
	IssuerPassword     = "password"
	IssuerPrivateToken = "private_token"
)

// 自定义非到处类型,避免外部包实例化

type tokenContextKey struct{}

var (
	CtxTokenKey = tokenContextKey{}
)

var (
	CookieNotFound = exception.NewUnauthorized("cookie %s not found", AccessTokenCookieName)
)

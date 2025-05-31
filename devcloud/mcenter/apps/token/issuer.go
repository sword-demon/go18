// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package token

import (
	"context"
	"fmt"
	"math/rand/v2"
)

var (
	charSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// MakeBearer 产生随机字符串
func MakeBearer(length int) string {

	t := make([]byte, 0)
	for range length {
		index := rand.IntN(len(charSet)) // 直接调用，无需播种
		t = append(t, charSet[index])
	}

	return string(t)
}

// issuers 颁发器的容器
var issuers = map[string]Issuer{}

// RegistryIssuer 注册颁发器
func RegistryIssuer(name string, p Issuer) {
	issuers[name] = p
}

// Issuer 颁发器必须实现的接口
type Issuer interface {
	// IssueToken 颁发 token 的接口
	IssueToken(context.Context, IssueParameter) (*Token, error)
}

func GetIssuer(name string) Issuer {
	fmt.Println(issuers)
	return issuers[name]
}

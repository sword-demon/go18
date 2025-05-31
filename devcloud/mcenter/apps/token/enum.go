// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package token

// SOURCE 来源定义
type SOURCE int

const (
	// SourceUnknown 未知
	SourceUnknown SOURCE = iota
	// SourceWeb web
	SourceWeb
	// SourceIos ios
	SourceIos
	// SourceAndroid android
	SourceAndroid
	// SourcePc PC
	SourcePc
	// SourceApi api 调用
	SourceApi SOURCE = 10
)

type LockType int

const (
	// LockTypeRevoke 用户退出登录
	LockTypeRevoke LockType = iota
	// LockTypeTokenExpired 刷新 token 过期,会话中断
	LockTypeTokenExpired
	// LockTypeOtherPlaceLoggedIn 异地登录
	LockTypeOtherPlaceLoggedIn
	// LockTypeOtherIpLoggedIn 异常 ip 登录
	LockTypeOtherIpLoggedIn
)

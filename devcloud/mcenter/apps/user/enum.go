// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package user

type PROVIDER int32

const (
	// ProviderLocal 本地数据库
	ProviderLocal PROVIDER = 0
	// ProviderLdap 来源 LDAP
	ProviderLdap PROVIDER = 1
	// ProviderFeishu 来源飞书
	ProviderFeishu PROVIDER = 2
	// 来源钉钉
	ProviderDingding PROVIDER = 3
	// 来源企业微信
	ProviderWeichatWork PROVIDER = 4
)

type CreateType int

const (
	// CreateTypeInit 系统初始化
	CreateTypeInit CreateType = iota
	// CreateTypeAdmin 管理员创建
	CreateTypeAdmin
	// CreateTypeRegistry 用户自己注册
	CreateTypeRegistry
)

// TYPE 用户类型
type TYPE int32

const (
	TypeSub TYPE = 0
)

type SEX int

const (
	SexUnknown SEX = iota
	SexMale
	SexFemale
)

type DescribeBy int

const (
	DescribeByID DescribeBy = iota
	DescribeByUsername
)

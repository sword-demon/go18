// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package endpoint

type AccessMode int8

const (
	// AccessModeRead 只读模式
	AccessModeRead = iota
	// AccessModeReadWrite 读写模式
	AccessModeReadWrite
)

const (
	MetaRequiredAuthKey      = "required_auth"
	MetaRequiredCodeKey      = "required_code"
	MetaRequiredPermKey      = "required_perm"
	MetaRequiredRoleKey      = "required_role"
	MetaRequiredAuditKey     = "required_audit"
	MetaRequiredNamespaceKey = "required_namespace"
	MetaResourceKey          = "resource"
	MetaActionKey            = "action"
)

// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package comptroller

const (
	// MetaAuditKey 定义开启审计的元数据键
	MetaAuditKey = "audit"
)

func Enable(v bool) (string, bool) {
	return MetaAuditKey, v
}

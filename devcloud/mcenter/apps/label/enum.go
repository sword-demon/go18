// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package label

type ValueType string

const (
	ValueTypeText     ValueType = "text"
	ValueTypeBoolean  ValueType = "bool"
	ValueTypeEnum     ValueType = "enum"
	ValueTypeHttpEnum ValueType = "http_enum" // 基于 url 的远程选项拉取,仅存储 url 地址,前端自己处理
)

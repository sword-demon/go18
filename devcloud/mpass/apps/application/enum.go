// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package application

type Type int32

const (
	TypeSourceCode     Type = 0  // 源代码
	TypeContainerImage Type = 1  // 容器镜像
	TypeOther          Type = 15 // 其他类型
)

var (
	TypeName = map[Type]string{
		TypeSourceCode:     "SourceCode",
		TypeContainerImage: "ContainerImage",
		TypeOther:          "Other",
	}
)

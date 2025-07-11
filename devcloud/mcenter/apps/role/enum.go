// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package role

const (
	ADMIN = "admin"
)

type MatchBy int32

const (
	// MatchById 针对某一个具体的接口进行授权
	MatchById = iota
	// MatchByLabel 通过标签来进行 api 接口授权
	MatchByLabel
	// MatchByResourceAction 通过资源和动作来进行授权 user::list
	MatchByResourceAction
	// MatchByResourceAccessMode 通过资源的访问模式来进行授权
	MatchByResourceAccessMode
)

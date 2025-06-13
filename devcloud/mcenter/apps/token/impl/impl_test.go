// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

var (
	svc token.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSet()
	svc = token.GetService()
}

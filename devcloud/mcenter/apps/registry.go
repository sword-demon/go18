// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package apps

import (
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/user/api"
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/user/impl"

	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/token/api"
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/token/impl"

	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint/impl"
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/role/impl"

	// token 颁发器
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps/token/issuers"
	// 鉴权中间件
	_ "github.com/sword-demon/go18/devcloud/mcenter/permission"
)

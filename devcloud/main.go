// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package main

import (
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	// 加载业务对象
	// 里面所有的 init 执行
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps"
	// 加载 api doc
	_ "github.com/infraboard/mcube/v2/ioc/apps/apidoc/restful"
	// 健康检查
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/restful"
	// metric
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
)

func main() {
	cmd.Start()
}

package main

import (
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	// 加载业务对象
	// 里面所有的 init 执行
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps"
	// 加载 api doc
	_ "github.com/infraboard/mcube/v2/ioc/apps/apidoc"
	// 健康检查
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/restful"
	// metric
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
)

func main() {
	cmd.Start()
}

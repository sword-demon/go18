// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package test

import (
	"fmt"
	"github.com/infraboard/mcube/v2/ioc"
	"os"
	// 加载的业务对象
	_ "github.com/sword-demon/go18/devcloud/mcenter/apps"
)

// DevelopmentSet 测试验证
func DevelopmentSet() {
	// import 后自动执行的逻辑
	// 工具对象的初始化,需要的是配置文件的绝对路径
	configPath := os.Getenv("GOPATH") + "/src/github.com/sword-demon/go18/devcloud/etc/application.toml"
	fmt.Println(configPath)
	ioc.DevelopmentSetupWithPath(configPath)
}

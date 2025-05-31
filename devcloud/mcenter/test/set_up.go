package test

import (
	"fmt"
	"github.com/infraboard/mcube/v2/ioc"
	"os"
	// 加载的业务对象
	//_ "github.com/sword-demon/go18/devcloud/mcenter/apps"
)

// DevelopmentSet 测试验证
func DevelopmentSet() {
	// import 后自动执行的逻辑
	// 工具对象的初始化,需要的是配置文件的绝对路径
	configPath := os.Getenv("GOPATH") + "/src/github.com/sword-demon/go18/devcloud/etc/application.toml"
	fmt.Println(configPath)
	ioc.DevelopmentSetupWithPath(configPath)
}

package test

import (
	"github.com/infraboard/mcube/v2/ioc"
	_ "github.com/sword-demon/go18/devcloud/mpass/apps"
	"os"
)

func DevelopmentSetup() {
	configPath := os.Getenv("GOPATH") + "/src/github.com/sword-demon/go18/devcloud/etc/application.toml"
	ioc.DevelopmentSetupWithPath(configPath)
}

// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

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

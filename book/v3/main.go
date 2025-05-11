// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/go18/book/v2/config"
	"github.com/sword-demon/go18/book/v3/exception"
	"github.com/sword-demon/go18/book/v3/handlers"
	"log"
)

func main() {
	// 加载配置
	path := "application.yaml"
	err := config.LoadConfigFromYaml(path)
	if err != nil {
		log.Println(err)
		return
	}

	server := gin.New()
	server.Use(gin.Logger(), exception.Recovery())
	h := handlers.NewBookApiHandler()
	h.Registry(server)

	ac := config.C().Application
	if err := server.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)); err != nil {
		panic(err)
	}
}

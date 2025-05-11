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

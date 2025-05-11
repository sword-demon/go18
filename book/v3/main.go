package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/go18/book/v2/config"
	"github.com/sword-demon/go18/book/v3/handlers"
)

func main() {
	// 加载配置
	path := "application.yaml"
	err := config.LoadConfigFromYaml(path)
	if err != nil {
		return
	}

	r := gin.Default()
	h := handlers.NewBookApiHandler()
	h.Registry(r)

	ac := config.C().Application
	if err := r.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)); err != nil {
		panic(err)
	}
}

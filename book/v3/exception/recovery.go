// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		// 非业务异常
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%#v", err))
		c.Abort()
	})
}

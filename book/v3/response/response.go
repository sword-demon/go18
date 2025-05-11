// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/go18/book/v3/exception"
	"net/http"
)

func OK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()
}

func Failed(ctx *gin.Context, err error) {
	// 一种是业务异常
	if e, ok := err.(*exception.ApiException); ok {
		ctx.JSON(e.HttpCode, e)
		ctx.Abort()
		return
	}

	// 另一种是非业务异常
	ctx.JSON(http.StatusInternalServerError, exception.NewApiException(http.StatusInternalServerError, err.Error()))
	ctx.Abort()
}

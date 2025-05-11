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

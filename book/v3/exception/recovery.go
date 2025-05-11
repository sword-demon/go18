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

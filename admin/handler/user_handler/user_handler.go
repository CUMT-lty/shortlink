package user_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByUsername(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{ // 响应都是成功响应，但会返回错误码
		"code":    1,
		"message": "hi" + c.Param("username"),
	})
}

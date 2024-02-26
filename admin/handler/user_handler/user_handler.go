package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lty/shortlink/admin/common/convertion/errorcode"
	"github.com/lty/shortlink/admin/common/convertion/result"
	user_service "github.com/lty/shortlink/admin/service"
	"net/http"
)

func GetUserByUsername(c *gin.Context) {
	userResult := user_service.GetUserByUsername(c.Param("username"))
	if userResult.ID > 0 { // 用户存在
		c.AbortWithStatusJSON(http.StatusOK, result.GetSuccessResult(userResult))
	} else { // 用户不存在
		c.AbortWithStatusJSON(http.StatusOK, result.GetFailureResult(errorcode.UserNull))
	}
}

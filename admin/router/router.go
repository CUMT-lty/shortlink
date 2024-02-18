package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lty/shortlink/admin/handler/user_handler"
)

func Register() *gin.Engine {
	r := gin.Default()
	// 两类路由：1.用户操作 2. 用户信息脱敏
	userGrop := r.Group("/api/short-link/admin/v1/user")       // gin 路由分组
	userGrop.GET("/:username", user_handler.GetUserByUsername) // gin 动态路由
	return r
}

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
	userGrop.GET("/has-username", user_handler.HasUsername)    // 检查用户名是否存在
	userGrop.POST("", user_handler.Register)                   // 注册
	return r
}

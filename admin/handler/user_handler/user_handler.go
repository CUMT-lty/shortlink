package user_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lty/shortlink/admin/common/convertion/errorcode"
	"github.com/lty/shortlink/admin/common/convertion/result"
	"github.com/lty/shortlink/admin/dto/req"
	user_service "github.com/lty/shortlink/admin/service"
	"net/http"
)

// GetUserByUsername 通过用户名查询用户
func GetUserByUsername(c *gin.Context) {
	userResult := user_service.GetUserByUsername(c.Param("username")) // 动态路由参数匹配
	if userResult.ID > 0 {                                            // 用户存在
		c.AbortWithStatusJSON(http.StatusOK, result.GetSuccessResult(userResult))
	} else { // 用户不存在
		c.AbortWithStatusJSON(http.StatusOK, result.GetFailureResult(errorcode.UserNull))
	}
}

// TODO：用户中的 phone 字段优雅脱敏处理

// HasUsername 根据用户名查询用户是否存在
func HasUsername(c *gin.Context) {
	has := user_service.HasUsername(c.Query("username"))               // 获取 get 参数
	c.AbortWithStatusJSON(http.StatusOK, result.GetSuccessResult(has)) // 只看 data 字段就可以
}

func Register(c *gin.Context) {
	var userRegisterReqDTO req.UserRegisterReqDTO
	err := c.ShouldBindBodyWith(&userRegisterReqDTO, binding.JSON)
	if err != nil {
		fmt.Println("post 参数解析错误") // TODO：错误处理
	}
	if errCode := user_service.Register(userRegisterReqDTO); len(errCode) != 0 {
		c.AbortWithStatusJSON(http.StatusOK, result.GetFailureResult(errCode)) // 注册失败
	} else {
		c.AbortWithStatusJSON(http.StatusOK, result.GetSuccessResult("register success")) // TODO：注册成功后这里的 data 部分是不是需要返回一个 id
	}
}

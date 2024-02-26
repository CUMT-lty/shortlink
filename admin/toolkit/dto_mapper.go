package toolkit

import (
	"github.com/lty/shortlink/admin/dto/resp"
	"github.com/lty/shortlink/admin/model"
)

func UserMapper(user model.User) resp.UserRespDTO {
	return resp.UserRespDTO{
		ID:       user.ID,
		Username: user.Username,
		RealName: user.RealName,
		Phone:    user.Phone,
		Email:    user.Email,
	}
}

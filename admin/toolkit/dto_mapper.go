package toolkit

import (
	"github.com/lty/shortlink/admin/dto/req"
	"github.com/lty/shortlink/admin/dto/resp"
	"github.com/lty/shortlink/admin/model"
)

func UserMapperDTO(user model.User) resp.UserRespDTO {
	return resp.UserRespDTO{
		ID:       user.ID,
		Username: user.Username,
		RealName: user.RealName,
		Phone:    user.Phone,
		Email:    user.Email,
	}
}

func RegisterReqDTOMapperUser(registerParam req.UserRegisterReqDTO) model.User {
	return model.User{
		Username: registerParam.Username,
		Password: registerParam.Password,
		RealName: registerParam.RealName,
		Phone:    registerParam.Phone,
		Email:    registerParam.Email,
	}
}

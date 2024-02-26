package service

import (
	"github.com/lty/shortlink/admin/dto/resp"
	"github.com/lty/shortlink/admin/model"
	"github.com/lty/shortlink/admin/toolkit"
)

func GetUserByUsername(username string) resp.UserRespDTO {
	u := new(model.User) // TODO：这里只是为了标识使用哪个数据表，后期可以把临时对象用 sync.Pool 管理起来
	user := u.GetUserByUsername(username)
	return toolkit.UserMapper(user)
}

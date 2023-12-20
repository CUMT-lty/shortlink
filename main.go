package main

import (
	"fmt"
	"github.com/lty/shortlink/admin/model/user"
	"github.com/lty/shortlink/config"
	"github.com/lty/shortlink/db"
)

func main() {
	fmt.Println(config.Conf.Common.CommonMysql.DSN)
	name := db.DB.Name()
	fmt.Println(name)

	u := user.User{
		Username: "lty111",
		Password: "pwd111",
		RealName: "lyt",
		Phone:    "13182371876",
		Email:    "m13182371876@163.com",
	}
	u.AddUser()
}

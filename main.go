package main

import (
	"fmt"
	"github.com/lty/shortlink/config"
	"github.com/lty/shortlink/db"
)

func main() {
	fmt.Println(config.Conf.Common.CommonMysql.DSN)
	name := db.DB.Name()
	fmt.Println(name)
}

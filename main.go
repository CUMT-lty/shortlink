package main

import (
	"fmt"
	"github.com/lty/shortlink/admin/toolkit"
)

func main() {
	//fmt.Println(config.Conf.Common.CommonMysql.DSN)
	//name := db.DB.Name()
	//fmt.Println(name)

	filter := toolkit.RegisterBloomFilter
	filter.Add([]byte("lty111"))
	fmt.Println(filter.Exists([]byte("lty111")))
	fmt.Println(filter.Exists([]byte("nothing")))
}

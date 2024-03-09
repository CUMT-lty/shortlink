package db

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
)

var RDB *redis.Redis
var once sync.Once

func init() {
	once.Do(func() {
		var err error
		redisConf := redis.RedisConf{ // TODO: 配置后面要整理为统一加载
			Host: "127.0.0.1:6379",
			Type: "node", // TODO : 单点模式、哨兵模式、集群模式？
			Pass: "",
			Tls:  false,
		}
		RDB, err = redis.NewRedis(redisConf)
		if err != nil {
			fmt.Printf("获取 redis 连接出错：%v", err) // TODO：日志和打印错误要重新封装
		}
	})
}

package toolkit

import (
	"github.com/lty/shortlink/db"
	"github.com/zeromicro/go-zero/core/bloom"
	"sync"
)

// RegisterBloomFilter 注册模块使用的布隆过滤器，保存用户名
var RegisterBloomFilter *bloom.Filter
var once sync.Once

func init() {
	once.Do(func() { // 单例模式
		RegisterBloomFilter = GetRegisterBloomFilter()
	})
}

func GetRegisterBloomFilter() *bloom.Filter {
	filter := bloom.New(db.RDB, "userRegisterCachePenetrationBloomFilter", uint(100000000))
	return filter
}

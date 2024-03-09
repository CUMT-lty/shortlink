package service

import (
	"github.com/lty/shortlink/admin/common/constant"
	"github.com/lty/shortlink/admin/common/convertion/errorcode"
	"github.com/lty/shortlink/admin/dto/req"
	"github.com/lty/shortlink/admin/dto/resp"
	"github.com/lty/shortlink/admin/model"
	"github.com/lty/shortlink/admin/toolkit"
	"github.com/lty/shortlink/db"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func GetUserByUsername(username string) resp.UserRespDTO {
	u := new(model.User) // TODO：这里只是为了标识使用哪个数据表，后期可以把临时对象用 sync.Pool 管理起来
	user := u.GetUserByUsername(username)
	return toolkit.UserMapperDTO(user)
	// TODO: 改成去布隆过滤器里面取
}

func HasUsername(username string) bool {
	u := new(model.User) // TODO：这里只是为了标识使用哪个数据表，后期可以把临时对象用 sync.Pool 管理起来
	user := u.GetUserByUsername(username)
	return user.ID > 0 // true 表示用户名存在，false 表示用户名不存在
}

func Register(registerParam req.UserRegisterReqDTO) errorcode.UserErrorCode {
	if has := HasUsername(registerParam.Username); has {
		return errorcode.UserNameExist
	}
	var err error
	user := toolkit.RegisterReqDTOMapperUser(registerParam)
	// TODO：获取分布式锁
	dtbLock := redis.NewRedisLock(db.RDB, constant.LockUserRegisterKey+registerParam.Username)
	//dtbLock.SetExpire(10) // TODO：过期时间设置多久（分布式锁的解决方案）
	acquire, err := dtbLock.Acquire() // 尝试获取锁 TODO：这里可以引入重试 retry
	switch {
	case err != nil: // 处理错误
		return errorcode.UserSaveError
	case acquire: // 正常业务逻辑
		defer dtbLock.Release()                  // defer 释放锁
		if err = user.InsertUser(); err != nil { // 先写 mysql
			return errorcode.UserSaveError
		}
		if err = toolkit.RegisterBloomFilter.Add([]byte(user.Username)); err != nil { // 再写 redis
			return errorcode.UserSaveError
		}
		// 这里如果第二阶段失败：如果再次注册该用户名，mysql 给用户名字段设置了唯一，所以 mysql 的 insert 操作会失败
		//case !acquire: // 没有拿到锁 wait?
	}
	return ""
}

package constant

var (
	// LockUserRegisterKey 用户注册时使用的分布式锁
	LockUserRegisterKey = "short-link:lock_user-register:"

	// LockGroupCreateKey 创建分组时使用的分布式锁
	LockGroupCreateKey = "short-link:lock_group-create:%s"

	// UserLoginKey 用户登录的缓存标识
	UserLoginKey = "short-link:login:"
)

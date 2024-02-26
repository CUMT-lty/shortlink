package errorcode

type IErrorCode interface {
	Code() string
	Message() string
}

type BaseErrorCode string

func (bec BaseErrorCode) Code() string {
	return string(bec)
}

func (bec BaseErrorCode) Message() string {
	return BaseErrorCodeMsgMap[bec]
}

type UserErrorCode string

func (uec UserErrorCode) Code() string {
	return string(uec)
}

func (uec UserErrorCode) Message() string {
	return UserErrorCodeMsgMap[uec]
}

var ( // BaseErrorCode
	ClientError BaseErrorCode = "A000001" // 一级宏观错误码 客户端错误

	UserRegisterError             BaseErrorCode = "A000100" // 二级宏观错误码 用户注册错误
	UserNameVerifyError           BaseErrorCode = "A000110"
	UserNameExistError            BaseErrorCode = "A000111"
	UserNameSensitiveError        BaseErrorCode = "A000112"
	UserNameSpecialCharacterError BaseErrorCode = "A000113"
	PasswordVerifyError           BaseErrorCode = "A000120"
	PasswordShortError            BaseErrorCode = "A000121"
	PhoneVerifyError              BaseErrorCode = "A000151"

	IdempotentTokenNullError   BaseErrorCode = "A000200" // 二级宏观错误码 系统请求缺少幂等Token
	IdempotentTokenDeleteError BaseErrorCode = "A000201"

	ServiceError BaseErrorCode = "B000001" // 一级宏观错误码 系统执行出错

	ServiceTimeoutError BaseErrorCode = "B000100" // 二级宏观错误码 系统执行超时

	RemoteError BaseErrorCode = "C000001" // 一级宏观错误码 调用第三方服务出错
)

var BaseErrorCodeMsgMap = map[BaseErrorCode]string{

	ClientError: "用户端错误", // 一级宏观错误码 客户端错误

	UserRegisterError:             "用户注册错误", // 二级宏观错误码 用户注册错误
	UserNameVerifyError:           "用户名校验失败",
	UserNameExistError:            "用户名已存在",
	UserNameSensitiveError:        "用户名包含敏感词",
	UserNameSpecialCharacterError: "用户名包含特殊字符",
	PasswordVerifyError:           "密码校验失败",
	PasswordShortError:            "密码长度不够",
	PhoneVerifyError:              "手机格式校验失败",

	IdempotentTokenNullError:   "幂等Token为空", // 二级宏观错误码 系统请求缺少幂等Token
	IdempotentTokenDeleteError: "幂等Token已被使用或失效",

	ServiceError: "系统执行出错", // 一级宏观错误码 系统执行出错

	ServiceTimeoutError: "系统执行超时", // 二级宏观错误码 系统执行超时

	RemoteError: "调用第三方服务出错", // 一级宏观错误码 调用第三方服务出错
}

var ( // UserErrorCode
	UserNull      UserErrorCode = "B000200"
	UserNameExist UserErrorCode = "B000201"
	UserExist     UserErrorCode = "B000202"
	UserSaveError UserErrorCode = "B000203"
)

var UserErrorCodeMsgMap = map[UserErrorCode]string{
	UserNull:      "用户记录不存在",
	UserNameExist: "用户名已存在",
	UserExist:     "用户记录已存在",
	UserSaveError: "用户记录新增失败",
}

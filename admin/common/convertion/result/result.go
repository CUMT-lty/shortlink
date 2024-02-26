package result

import (
	"github.com/lty/shortlink/admin/common/convertion/errorcode"
	"github.com/lty/shortlink/project/common/convertion/result"
)

const SUCCESS_CODE = "0"

type Result struct { // TODO: 分布式 ID 的解决方案
	//serialVersionUID int64 // TODO: 这个是干嘛的
	Code      string      // 返回码
	Message   string      // 返回消息
	Data      interface{} // 响应数据
	RequestId string      // TODO: 请求 id
	Success   bool        // 是否成功
}

func GetSuccessResult(data any) result.Result {
	return result.Result{
		Code:    result.SUCCESS_CODE,
		Data:    data,
		Success: true,
	}
}

func GetFailureResult(errorCode errorcode.IErrorCode) result.Result {
	return result.Result{
		Code:    errorCode.Code(),
		Message: errorCode.Message(),
		Success: false,
	}
}

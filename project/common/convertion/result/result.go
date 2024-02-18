package result

const SUCCESS_CODE = "0"

type Result struct {
	//serialVersionUID int64 // TODO: 这个是干嘛的
	Code      string      // 返回码
	Message   string      // 返回消息
	Data      interface{} // 响应数据
	RequestId string      // TODO: 请求 id
	Success   bool        // 是否成功
}

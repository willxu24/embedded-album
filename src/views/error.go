package views

const (
	// 基本响应
	Success     = 20000
	ErrorClient = 40000
	Error       = 50000

	// 进阶响应
	ErrorInput   = 40001
	ErrorAuthMsg = 40011
)

var errMap = map[int]string{
	// 基本响应
	Success: "",
	Error:   "服务端错误",

	// 进阶响应
	ErrorInput:   "客户端输入错误",
	ErrorAuthMsg: "用户名或者密码不正确",
}

func getErr(code int) string {
	msg, err := errMap[code]
	if !err {
		return "未知错误"
	}
	return msg
}

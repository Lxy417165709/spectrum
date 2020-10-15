package model
type Response struct {
	Err  string      `json:"err"`  // 错误消息，如格式错误
	Msg  string      `json:"msg"`  // 一般消息，如登录成功
	Data interface{} `json:"data"` // 响应数据
}

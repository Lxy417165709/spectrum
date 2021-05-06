package model

type Response struct {
	Err  string      `json:"err"`  // 错误消息，如格式错误
	Msg  string      `json:"msg"`  // 一般消息，如登录成功
	Data interface{} `json:"data"` // 响应数据
}

type Request struct {
	Object     string                 `json:"object"`		// 请求RPC对象
	Function   string                 `json:"function"`		// 请求RPC对象的函数
	Parameters map[string]interface{} `json:"parameters"`	// 参数
}

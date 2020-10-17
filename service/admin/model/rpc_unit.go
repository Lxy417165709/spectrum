package model


type RpcUnit struct{
	ResFunc func()interface{}
	ReqFunc func()interface{}
	SuccessMsg string
	FailMsg string
}


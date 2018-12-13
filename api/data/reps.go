package data

//Resp response
type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

//NewFineResp ok response 200
func NewFineResp(msg string, data interface{}) Resp {
	return Resp{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

//NewErrorResp error response 400
func NewErrorResp(data interface{}) Resp {
	return Resp{
		Code: 400,
		Msg:  "请求错误～",
		Data: data,
	}
}

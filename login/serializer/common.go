package serializer

/*
	code:0,
	msg:创建成,
	data:
	error:
*/
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   interface{} `json:"msg"`
	Error string      `json:"error"`
}

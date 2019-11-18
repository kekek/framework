package response

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"info"`
	Data interface{} `json:"data"`
}

func ActionResponse(actionCode int, data interface{}) Response {
	return Response{Code: actionCode, Msg: "ok", Data: data}
}

func OKResponse(data interface{}) Response {
	return Response{Code: 200, Msg: "ok", Data: data}
}

func ErrorResponse(code int, msg interface{}) Response {
	return Response{Code: code, Msg: msg, Data: nil}
}

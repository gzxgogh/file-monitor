package main

type SendPaoPaoMsgReq struct {
	Msg      string `json:"msg"`
	GroupId  string `json:"groupId"`
	ToCustId string `json:"toCustId"`
}

type CreateMonitorReq struct {
	Path string `json:"path" form:"path"`
}

type Result struct {
	Status int         `json:"status" bson:"status"`
	Msg    string      `json:"msg" bson:"msg"`
	Data   interface{} `json:"data" bson:"data"`
}

func Error(code int, msg string) Result {
	return Result{
		Status: code,
		Msg:    msg,
		Data:   nil,
	}
}

func Success(data interface{}) Result {
	return Result{
		Status: 200,
		Msg:    "success",
		Data:   data,
	}
}

package response

import "time"

type ApiResp struct {
	ErrCode   int         `json:"errcode"`
	ErrMsg    string      `json:"errmsg"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

func Success(value interface{}) ApiResp {
	return ApiResp{
		ErrCode:   0,
		ErrMsg:    "success",
		Data:      value,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func Fail(code int, msg string) ApiResp {
	return ApiResp{
		ErrCode:   code,
		ErrMsg:    msg,
		Data:      nil,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}

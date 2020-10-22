package utils

import (
	"encoding/json"
)

type Result struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type tmp struct {
	Item  interface{} `json:"item,omitempty"`
	Count int         `json:"count,omitempty"`
}

const (
	Success int = 200
	Fail    int = 300
)

func NewResult(code int, msg string, data ...interface{}) []byte {
	if len(data) > 0 {
		res := Result{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
		buf, _ := json.Marshal(res)
		return buf
	} else {
		res := Result{
			Code: code,
			Msg:  msg,
		}
		buf, _ := json.Marshal(res)
		return buf
	}
}

func NewPage(code int, msg string, data interface{}, count int) Result {
	res := Result{
		Code: code,
		Msg:  msg,
		Data: tmp{
			Item:  data,
			Count: count,
		},
	}
	return res
}

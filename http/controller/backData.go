package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Return struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HTTPSuccess(w http.ResponseWriter, data interface{}) {
	marshal, _ := json.Marshal(Return{
		Code: 1,
		Msg:  "成功",
		Data: data,
	})
	_, err := fmt.Fprintf(w, string(marshal))
	if err != nil {
		return
	}
}

func HTTPError(w http.ResponseWriter, code int, msg string) {
	marshal, _ := json.Marshal(Return{
		Code: code,
		Msg:  msg,
	})
	_, err := fmt.Fprintf(w, string(marshal))
	if err != nil {
		return
	}
}

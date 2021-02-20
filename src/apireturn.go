package main

import (
	"encoding/json"
	"net/http"
)

type ApiReturn struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func returnOk(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(&ApiReturn{
		Retcode: 0,
		Message: "OK",
	})
}
func returnErr(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(&ApiReturn{
		Retcode: -1,
		Message: "请求出错",
	})
}

func returnErrMsg(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(&ApiReturn{
		Retcode: -1,
		Message: msg,
	})
}

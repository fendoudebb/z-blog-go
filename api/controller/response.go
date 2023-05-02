package controller

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(w http.ResponseWriter, data any) {
	payload, _ := json.Marshal(response{
		Data: data,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

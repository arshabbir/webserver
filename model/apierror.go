package model

type ApiError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

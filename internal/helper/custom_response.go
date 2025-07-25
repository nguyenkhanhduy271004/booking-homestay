package helper

import "strings"

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data any) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data any) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

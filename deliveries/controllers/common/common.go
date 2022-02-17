package common

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func InternalServerError() Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Message: "terdapat error di dalam server",
		Data:    nil,
	}
}

func BadRequest() Response {
	return Response{
		Code:    http.StatusBadRequest,
		Message: "input client tidak sesuai",
		Data:    nil,
	}
}

func UnAuthorized() Response {
	return Response{
		Code:    http.StatusUnauthorized,
		Message: "email atau kata sandi tidak valid",
		Data:    nil,
	}
}

func NotAcceptable() Response {
	return Response{
		Code:    http.StatusNotAcceptable,
		Message: "tidak dapat memproses nilai yang didapatkan",
		Data:    nil,
	}
}

func Success(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

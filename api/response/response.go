package response

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResOK(message string, data interface{}) Response {
	res := Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
	return res
}

func ResErr(status int, message string) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    map[string]interface{}{},
	}
	return res
}

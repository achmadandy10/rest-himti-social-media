package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(status int, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  status,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}

	return res
}

func ValidationErrorResponse(status int, message string, err interface{}, data interface{}) Response {
	errMessages := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errMessages[strings.ToLower(e.Field())] = e.ActualTag()
	}

	res := Response{
		Status:  status,
		Message: message,
		Error:   errMessages,
		Data:    data,
	}

	return res
}

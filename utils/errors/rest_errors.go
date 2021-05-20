package errors

import (
	"net/http"
)

type RestErr struct {
	Meassage string `json:"message"`
	Status   int    `json:"status"`
	Error    string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Meassage: message,
		Status:   http.StatusBadRequest,
		Error:    "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Meassage: message,
		Status:   http.StatusNotFound,
		Error:    "not_found",
	}
}

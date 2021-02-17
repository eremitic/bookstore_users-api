package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadReqErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "bad request",
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}

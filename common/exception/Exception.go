package exception

import "net/http"

type Exception struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}

func BadRequestException(message string, err error) *Exception {
	return &Exception{
		Message: message,
		Error:   err.Error(),
		Code:    http.StatusBadRequest,
	}
}

func InternalServerErrorException(message string, err error) *Exception {
	return &Exception{
		Message: message,
		Error:   err.Error(),
		Code:    http.StatusInternalServerError,
	}
}

func NotFoundException(message string, err error) *Exception {
	return &Exception{
		Message: message,
		Error:   err.Error(),
		Code:    http.StatusNotFound,
	}
}

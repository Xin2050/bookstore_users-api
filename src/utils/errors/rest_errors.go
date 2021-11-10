package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}
func NotFoundError(message string) *RestError {
	return &RestError{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not_found_error",
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "internal_server_error",
	}
}

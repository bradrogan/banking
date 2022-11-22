package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

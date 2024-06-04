package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error has occurred",
	}
}

func NewBadRequestError() error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: "bad request",
	}
}

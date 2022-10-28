package errutil

import (
	"errors"
	"net/http"
)

type InternalServerError struct {
	message  string
	httpCode int
}

func (i *InternalServerError) Error() string {
	return i.message
}

func (i *InternalServerError) GetHTTPCode() int {
	return i.httpCode
}

func NewInternalServerError(msg string) error {
	return &InternalServerError{
		message:  msg,
		httpCode: http.StatusInternalServerError,
	}
}

func IsInternalServerError(err error) bool {
	var expectedErr *InternalServerError
	return errors.As(err, &expectedErr)
}

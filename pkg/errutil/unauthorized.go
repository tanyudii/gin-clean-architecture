package errutil

import (
	"errors"
	"net/http"
)

type UnauthorizedError struct {
	message  string
	httpCode int
}

func (i *UnauthorizedError) Error() string {
	return i.message
}

func (i *UnauthorizedError) GetHTTPCode() int {
	return i.httpCode
}

func NewUnauthorizedError(msg string) error {
	return &UnauthorizedError{
		message:  msg,
		httpCode: http.StatusUnauthorized,
	}
}

func IsUnauthorizedError(err error) bool {
	var expectedErr *UnauthorizedError
	return errors.As(err, &expectedErr)
}

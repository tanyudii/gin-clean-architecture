package errutil

import (
	"errors"
	"net/http"
)

type UnauthenticatedError struct {
	message  string
	httpCode int
}

func (i *UnauthenticatedError) Error() string {
	return i.message
}

func (i *UnauthenticatedError) GetHTTPCode() int {
	return i.httpCode
}

func NewUnauthenticatedError(msg string) error {
	return &UnauthenticatedError{
		message:  msg,
		httpCode: http.StatusUnauthorized,
	}
}

func IsUnauthenticatedError(err error) bool {
	var expectedErr *UnauthenticatedError
	return errors.As(err, &expectedErr)
}

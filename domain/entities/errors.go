package entities

import (
	"github.com/vodeacloud/hr-api/pkg/errutil"
)

var (
	ErrUnsupportedGrantType      = errutil.NewBadRequestError("unsupported grant type")
	ErrInvalidCredentials        = errutil.NewBadRequestError("credentials doesn't match in our records")
	ErrInvalidClient             = errutil.NewBadRequestError("client is not registered")
	ErrPrivateOrPublicKeyInvalid = errutil.NewInternalServerError("private or public key is invalid")
	ErrInvalidOrExpiredAccess    = errutil.NewBadRequestError("access token is invalid or expired")
	ErrInvalidOrExpiredRefresh   = errutil.NewBadRequestError("refresh token is invalid or expired")
)

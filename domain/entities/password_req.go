package entities

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/pkg/errutil"
)

type PasswordEmailResponse struct {
	Message string `json:"message"`
}

type PasswordEmailRequest struct {
	Email string `form:"email"`
	URL   string `form:"url"`
}

func (r *PasswordEmailRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.Email == "" {
		fields["email"] = "email field is required"
	} else if !govalidator.IsEmail(r.Email) {
		fields["email"] = "email format invalid"
	}
	if r.URL == "" {
		fields["url"] = "url field is required"
	} else if !govalidator.IsURL(r.URL) {
		fields["url"] = "url is invalid"
	}
	return errutil.BadRequestOrNil(fields)
}

type PasswordResetResponse struct {
	Message string `json:"message"`
}

type PasswordResetRequest struct {
	Email                string `form:"email"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"passwordConfirmation"`
	Token                string `form:"token"`
}

func (r *PasswordResetRequest) Validate() error {
	cfg := config.GetConfig()
	fields := errutil.ErrorField{}
	if r.Email == "" {
		fields["email"] = "email field is required"
	} else if !govalidator.IsEmail(r.Email) {
		fields["email"] = "email format invalid"
	}
	if r.Password == "" {
		fields["password"] = "password field is required"
	} else if len(r.Password) < cfg.UserPasswordMinLength {
		fields["password"] = fmt.Sprintf("password minimum %d characters", cfg.UserPasswordMinLength)
	} else if r.Password != r.PasswordConfirmation {
		fields["passwordConfirmation"] = "password confirmation doesn't match"
	}
	if r.Token == "" {
		fields["token"] = "token field is required"
	}
	return errutil.BadRequestOrNil(fields)
}

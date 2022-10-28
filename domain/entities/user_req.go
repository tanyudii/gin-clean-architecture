package entities

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"github.com/vodeacloud/hr-api/pkg/pagination"
)

type UsersResponse struct {
	Users      []*User                `json:"data"`
	Pagination *pagination.Pagination `json:"meta"`
}

type UsersResponseWithoutPagination struct {
	Users []*User `json:"data"`
}

type UserResponse struct {
	User *User `json:"data"`
}

type GetUsersByQueryRequest struct {
	Page   int32  `form:"page"`
	Limit  int32  `form:"limit"`
	Sort   int32  `form:"sort"`
	Search string `form:"search"`
}

func (r *GetUsersByQueryRequest) Validate() error {
	if r.Page <= 0 {
		return errutil.NewBadRequestError("page must be start from 1")
	} else if r.Limit <= 0 {
		return errutil.NewBadRequestError("limit must be start from 1")
	}
	return nil
}

type GetUserByIDRequest struct {
	ID int64 `uri:"id"`
}

func (r *GetUserByIDRequest) Validate() error {
	if r.ID <= 0 {
		return errutil.NewBadRequestError("id is invalid")
	}
	return nil
}

type CreateUserRequest struct {
	Name       string  `form:"name"`
	Email      string  `form:"email"`
	Password   string  `form:"password"`
	Phone      *string `form:"phone"`
	AvatarURL  *string `form:"avatarURL"`
	Type       *UserType
	UserDetail *UserDetailRequest `form:"userDetail"`
}

func (r *CreateUserRequest) Validate() error {
	r.SetEmptyFieldToNil()
	cfg := config.GetConfig()
	fields := errutil.ErrorField{}
	if r.Name == "" {
		fields["name"] = "name field is mandatory"
	}
	if r.Email == "" {
		fields["email"] = "email field is mandatory"
	} else if !govalidator.IsEmail(r.Email) {
		fields["email"] = "email format is invalid"
	}
	if r.Password == "" {
		fields["password"] = "password field is mandatory"
	} else if len(r.Password) < cfg.UserPasswordMinLength {
		fields["password"] = fmt.Sprintf("password minimum %d characters", cfg.UserPasswordMinLength)
	}
	if r.AvatarURL != nil && !govalidator.IsURL(*r.AvatarURL) {
		fields["avatarURL"] = "avatar is invalid"
	}
	if r.Type != nil && !r.Type.Valid() {
		fields["type"] = "type is invalid"
	}
	if r.UserDetail == nil {
		fields["userDetail"] = "user detail field is mandatory"
	} else {
		r.UserDetail.Validate(fields, "userDetail.")
	}
	return errutil.BadRequestOrNil(fields)
}

func (r *CreateUserRequest) SetEmptyFieldToNil() {
	if r.Phone != nil && *r.Phone == "" {
		r.Phone = nil
	}
	if r.AvatarURL != nil && *r.AvatarURL == "" {
		r.AvatarURL = nil
	}
}

type UpdateUserRequest struct {
	ID         int64              `uri:"id"`
	Name       string             `form:"name"`
	Email      string             `form:"email"`
	Phone      *string            `form:"phone"`
	AvatarURL  *string            `form:"avatarURL"`
	UserDetail *UserDetailRequest `form:"userDetail"`
}

func (r *UpdateUserRequest) Validate() error {
	r.SetEmptyFieldToNil()
	fields := errutil.ErrorField{}
	if r.ID <= 0 {
		fields["id"] = "id is invalid"
	}
	if r.Name == "" {
		fields["name"] = "name field is mandatory"
	}
	if r.Email == "" {
		fields["email"] = "email field is mandatory"
	} else if !govalidator.IsEmail(r.Email) {
		fields["email"] = "email format is invalid"
	}
	if r.AvatarURL != nil && !govalidator.IsURL(*r.AvatarURL) {
		fields["avatarURL"] = "avatar is invalid"
	}
	if r.UserDetail == nil {
		fields["userDetail"] = "user detail field is mandatory"
	} else {
		r.UserDetail.Validate(fields, "userDetail.")
	}
	return errutil.BadRequestOrNil(fields)
}

func (r *UpdateUserRequest) SetEmptyFieldToNil() {
	if r.Phone != nil && *r.Phone == "" {
		r.Phone = nil
	}
	if r.AvatarURL != nil && *r.AvatarURL == "" {
		r.AvatarURL = nil
	}
}

type DeleteUserRequest struct {
	ID int64 `uri:"id"`
}

func (r *DeleteUserRequest) Validate() error {
	if r.ID <= 0 {
		return errutil.NewBadRequestError("id is invalid")
	}
	return nil
}

type DeleteUserBulkRequest struct {
	IDs []int64 `form:"ids"`
}

func (r *DeleteUserBulkRequest) Validate() error {
	if len(r.IDs) == 0 {
		return errutil.NewBadRequestError("ids is required")
	}
	return nil
}

type UserDetailRequest struct {
	CompanySerial string `form:"companySerial"`
}

func (r *UserDetailRequest) Validate(fields errutil.ErrorField, prefix string) {
	if r.CompanySerial == "" {
		fields[prefix+"companySerial"] = "company serial field is required"
	}
}

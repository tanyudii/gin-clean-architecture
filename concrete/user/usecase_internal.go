package user

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/pkg/errutil"
)

func (u *Usecase) validateUniqueUser(ctx context.Context, user *entities.User) error {
	userByEmail, err := u.userRepo.GetUserByEmail(ctx, user.Name)
	if err != nil && !errutil.IsNotFoundError(err) {
		return err
	}
	if userByEmail != nil && user.ID != userByEmail.ID {
		return errutil.NewBadRequestError("email has already been registered")
	}

	if user.Phone != nil {
		userByPhone, err := u.userRepo.GetUserByPhone(ctx, *user.Phone)
		if err != nil && !errutil.IsNotFoundError(err) {
			return err
		}
		if userByPhone != nil && user.ID != userByPhone.ID {
			return errutil.NewBadRequestError("phone has already been registered")
		}
	}

	return nil
}

func (u *Usecase) validateUserDetail(ctx context.Context, detail *entities.UserDetail) error {
	if detail == nil {
		return errutil.NewBadRequestError("user detail is empty")
	}

	companyBySerial, err := u.companyRepo.GetCompanyBySerial(ctx, detail.CompanySerial)
	if err != nil && !errutil.IsNotFoundError(err) {
		return err
	}
	if companyBySerial == nil {
		return errutil.NewBadRequestError("company is invalid")
	}
	return nil
}

func (u *Usecase) appendUserRelations(ctx context.Context, users []*entities.User) error {
	mapUserBySerial := map[string]*entities.User{}
	var userSerials []string
	for _, user := range users {
		userSerials = append(userSerials, user.Serial)
		mapUserBySerial[user.Serial] = user
	}

	mapUserDetailByUserSerial, err := u.userRepo.GetMapUserDetailByUserSerial(ctx, userSerials)
	if err != nil {
		return err
	}

	for _, user := range users {
		if userDetail, ok := mapUserDetailByUserSerial[user.Serial]; ok {
			user.UserDetail = userDetail
		}
	}

	return nil
}

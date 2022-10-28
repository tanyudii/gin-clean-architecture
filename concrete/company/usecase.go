package company

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	companyRepo repositories.CompanyRepository
	userUc      usecases.UserUsecase
}

func NewUsecase(
	companyRepo repositories.CompanyRepository,
	userUc usecases.UserUsecase,
) usecases.CompanyUsecase {
	return &Usecase{
		companyRepo: companyRepo,
		userUc:      userUc,
	}
}

func (u *Usecase) RegisterCompany(ctx context.Context, r *entities.RegisterCompanyRequest) error {
	if err := r.Validate(); err != nil {
		return err
	}

	company := &entities.Company{Serial: uuid.NewString(), Name: r.Name}
	if err := u.companyRepo.CreateCompany(ctx, company); err != nil {
		return err
	}

	userTypeAdmin := entities.UserTypeAdmin
	user := &entities.CreateUserRequest{
		Name:     fmt.Sprintf("Admin %s", r.Name),
		Email:    r.AdminEmail,
		Password: r.AdminPassword,
		Type:     &userTypeAdmin,
		UserDetail: &entities.UserDetailRequest{
			CompanySerial: company.Serial,
		},
	}

	_, err := u.userUc.CreateUser(ctx, user)

	return err
}

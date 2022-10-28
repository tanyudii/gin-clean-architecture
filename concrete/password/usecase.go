package password

import (
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	passwordRepo repositories.PasswordRepository
	userRepo     repositories.UserRepository
	emailUc      usecases.EmailUsecase
}

func NewUsecase(
	passwordRepo repositories.PasswordRepository,
	userRepo repositories.UserRepository,
	emailUc usecases.EmailUsecase,
) usecases.PasswordUsecase {
	return &Usecase{
		passwordRepo: passwordRepo,
		userRepo:     userRepo,
		emailUc:      emailUc,
	}
}

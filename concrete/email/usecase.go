package email

import (
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	emailTemplateRepo repositories.EmailTemplateRepository
}

func NewUsecase(
	emailTemplateRepo repositories.EmailTemplateRepository,
) usecases.EmailUsecase {
	return &Usecase{
		emailTemplateRepo: emailTemplateRepo,
	}
}

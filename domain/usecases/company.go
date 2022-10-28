package usecases

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type CompanyUsecase interface {
	RegisterCompany(ctx context.Context, r *entities.RegisterCompanyRequest) error
}

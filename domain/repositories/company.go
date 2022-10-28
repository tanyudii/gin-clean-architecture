package repositories

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type CompanyRepository interface {
	GetCompanyBySerial(ctx context.Context, serial string) (*entities.Company, error)
	CreateCompany(ctx context.Context, company *entities.Company) error
	UpdateCompany(ctx context.Context, company *entities.Company) error
	UpsertCompanyDetail(ctx context.Context, companyDetail *entities.CompanyDetail) error
}

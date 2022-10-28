package company

import (
	"context"
	"errors"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) repositories.CompanyRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetCompanyBySerial(_ context.Context, serial string) (*entities.Company, error) {
	var role entities.Company
	if err := r.db.Where("companies.serial = ?", serial).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("company serial not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *Repository) CreateCompany(_ context.Context, company *entities.Company) error {
	return r.db.Create(&company).Error
}

func (r *Repository) UpdateCompany(_ context.Context, company *entities.Company) error {
	return r.db.Save(&company).Error
}

func (r *Repository) UpsertCompanyDetail(_ context.Context, companyDetail *entities.CompanyDetail) error {
	return r.db.Save(companyDetail).Error
}

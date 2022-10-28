package permission

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) repositories.PermissionRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllPermissions(_ context.Context) ([]*entities.Permission, error) {
	var permissions []*entities.Permission
	if err := r.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *Repository) GetPermissionsByCodes(_ context.Context, codes []string) ([]*entities.Permission, error) {
	//skip immediately when codes are empty
	if len(codes) == 0 {
		return nil, nil
	}

	var permissions []*entities.Permission
	if err := r.db.Where("permissions.code IN (?)", codes).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *Repository) GetMapPermissionByCode(ctx context.Context, codes []string) (map[string]*entities.Permission, error) {
	permissions, err := r.GetPermissionsByCodes(ctx, codes)
	if err != nil {
		return nil, err
	}

	mapPermissionByCode := map[string]*entities.Permission{}
	for _, permission := range permissions {
		mapPermissionByCode[permission.Code] = permission
	}

	return mapPermissionByCode, nil
}

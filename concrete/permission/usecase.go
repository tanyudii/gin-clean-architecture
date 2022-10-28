package permission

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	permissionRepo repositories.PermissionRepository
}

func NewUsecase(
	permissionRepo repositories.PermissionRepository,
) usecases.PermissionUsecase {
	return &Usecase{
		permissionRepo: permissionRepo,
	}
}

func (u Usecase) GetAllPermissions(ctx context.Context) (*entities.PermissionsResponseWithoutPagination, error) {
	permissions, err := u.permissionRepo.GetAllPermissions(ctx)
	if err != nil {
		return nil, err
	}

	return &entities.PermissionsResponseWithoutPagination{
		Permissions: permissions,
	}, nil
}

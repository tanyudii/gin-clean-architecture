package usecases

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type PermissionUsecase interface {
	GetAllPermissions(ctx context.Context) (*entities.PermissionsResponseWithoutPagination, error)
}

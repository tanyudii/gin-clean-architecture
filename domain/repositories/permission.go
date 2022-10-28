package repositories

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type PermissionRepository interface {
	GetAllPermissions(ctx context.Context) ([]*entities.Permission, error)
	GetPermissionsByCodes(ctx context.Context, codes []string) ([]*entities.Permission, error)
	GetMapPermissionByCode(ctx context.Context, codes []string) (map[string]*entities.Permission, error)
}

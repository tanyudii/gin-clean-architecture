package repositories

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type RoleRepository interface {
	GetRolesByQuery(ctx context.Context, q *entities.RolesQuery) ([]*entities.Role, error)
	GetRolesBySerials(ctx context.Context, serials []string) ([]*entities.Role, error)
	GetRoleByID(ctx context.Context, id int64) (*entities.Role, error)
	GetRoleBySerial(ctx context.Context, serial string) (*entities.Role, error)
	GetRoleByName(ctx context.Context, name string) (*entities.Role, error)
	CreateRole(ctx context.Context, role *entities.Role) error
	UpdateRole(ctx context.Context, role *entities.Role) error
	DeleteRole(ctx context.Context, id int64) (*entities.Role, error)
	DeleteRoleBulk(ctx context.Context, ids []int64) ([]*entities.Role, error)
	GetRolePermissionsByRoleSerials(ctx context.Context, roleSerials []string) ([]*entities.RolePermission, error)
	SyncRolePermission(ctx context.Context, role *entities.Role, rolePermissions []*entities.RolePermission) error
}

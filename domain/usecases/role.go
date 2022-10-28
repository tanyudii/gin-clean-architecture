package usecases

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type RoleUsecase interface {
	GetRolesByQuery(ctx context.Context, r *entities.GetRolesByQueryRequest) (*entities.RolesResponse, error)
	GetRoleByID(ctx context.Context, r *entities.GetRoleByIDRequest) (*entities.RoleResponse, error)
	CreateRole(ctx context.Context, r *entities.CreateRoleRequest) (*entities.RoleResponse, error)
	UpdateRole(ctx context.Context, r *entities.UpdateRoleRequest) (*entities.RoleResponse, error)
	DeleteRole(ctx context.Context, r *entities.DeleteRoleRequest) (*entities.RoleResponse, error)
	DeleteRoleBulk(ctx context.Context, r *entities.DeleteRoleBulkRequest) (*entities.RolesResponseWithoutPagination, error)
}

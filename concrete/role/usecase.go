package role

import (
	"context"
	"github.com/google/uuid"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
	"github.com/vodeacloud/hr-api/pkg/pagination"
)

type Usecase struct {
	roleRepo       repositories.RoleRepository
	permissionRepo repositories.PermissionRepository
}

func NewUsecase(
	roleRepo repositories.RoleRepository,
	permissionRepo repositories.PermissionRepository,
) usecases.RoleUsecase {
	return &Usecase{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
	}
}

func (u *Usecase) GetRolesByQuery(ctx context.Context, r *entities.GetRolesByQueryRequest) (*entities.RolesResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	q := &entities.RolesQuery{
		Search:     r.Search,
		Sort:       r.Sort,
		Pagination: pagination.NewPagination(r.Page, r.Limit),
	}

	roles, err := u.roleRepo.GetRolesByQuery(ctx, q)
	if err != nil {
		return nil, err
	}

	if err = u.appendRoleRelations(ctx, roles); err != nil {
		return nil, err
	}

	return &entities.RolesResponse{
		Roles:      roles,
		Pagination: q.Pagination,
	}, nil
}

func (u *Usecase) GetRoleByID(ctx context.Context, r *entities.GetRoleByIDRequest) (*entities.RoleResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	role, err := u.roleRepo.GetRoleByID(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	if err = u.appendRoleRelations(ctx, []*entities.Role{role}); err != nil {
		return nil, err
	}

	return &entities.RoleResponse{
		Role: role,
	}, nil
}

func (u *Usecase) CreateRole(ctx context.Context, r *entities.CreateRoleRequest) (*entities.RoleResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	role := &entities.Role{
		Serial:       uuid.NewString(),
		Name:         r.Name,
		IsActive:     r.IsActive,
		ParentSerial: r.ParentSerial,
	}

	if err := u.validateUniqueRole(ctx, role); err != nil {
		return nil, err
	}

	if err := u.validateParentSerial(ctx, role); err != nil {
		return nil, err
	}

	rolePermissions, err := u.validatePermissionCodes(ctx, role, r.PermissionCodes)
	if err != nil {
		return nil, err
	}

	if err = u.roleRepo.CreateRole(ctx, role); err != nil {
		return nil, err
	}

	if err = u.roleRepo.SyncRolePermission(ctx, role, rolePermissions); err != nil {
		return nil, err
	}

	if err = u.appendRoleRelations(ctx, []*entities.Role{role}); err != nil {
		return nil, err
	}

	return &entities.RoleResponse{
		Role: role,
	}, nil
}

func (u *Usecase) UpdateRole(ctx context.Context, r *entities.UpdateRoleRequest) (*entities.RoleResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	role, err := u.roleRepo.GetRoleByID(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	role.Name = r.Name
	role.IsActive = r.IsActive
	role.ParentSerial = r.ParentSerial

	if err = u.validateUniqueRole(ctx, role); err != nil {
		return nil, err
	}

	if err = u.validateParentSerial(ctx, role); err != nil {
		return nil, err
	}

	rolePermissions, err := u.validatePermissionCodes(ctx, role, r.PermissionCodes)
	if err != nil {
		return nil, err
	}

	if err = u.roleRepo.UpdateRole(ctx, role); err != nil {
		return nil, err
	}

	if err = u.roleRepo.SyncRolePermission(ctx, role, rolePermissions); err != nil {
		return nil, err
	}

	if err = u.appendRoleRelations(ctx, []*entities.Role{role}); err != nil {
		return nil, err
	}

	return &entities.RoleResponse{
		Role: role,
	}, nil
}

func (u *Usecase) DeleteRole(ctx context.Context, r *entities.DeleteRoleRequest) (*entities.RoleResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	role, err := u.roleRepo.DeleteRole(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	return &entities.RoleResponse{
		Role: role,
	}, nil
}

func (u *Usecase) DeleteRoleBulk(ctx context.Context, r *entities.DeleteRoleBulkRequest) (*entities.RolesResponseWithoutPagination, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	roles, err := u.roleRepo.DeleteRoleBulk(ctx, r.IDs)
	if err != nil {
		return nil, err
	}

	return &entities.RolesResponseWithoutPagination{
		Roles: roles,
	}, nil
}

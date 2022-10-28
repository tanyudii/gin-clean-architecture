package role

import (
	"context"
	"fmt"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/pkg/errutil"
)

func (u *Usecase) validateUniqueRole(ctx context.Context, role *entities.Role) error {
	roleByName, err := u.roleRepo.GetRoleByName(ctx, role.Name)
	if err != nil && !errutil.IsNotFoundError(err) {
		return err
	}
	if roleByName != nil && role.ID != roleByName.ID {
		return errutil.NewBadRequestError("name has already been registered")
	}
	return nil
}

func (u *Usecase) validateParentSerial(ctx context.Context, role *entities.Role) error {
	if role.ParentSerial == nil {
		return nil
	}

	if role.Serial == *role.ParentSerial {
		return errutil.NewBadRequestError("supervisor/parent can't point to myself")
	}

	parentBySerial, err := u.roleRepo.GetRoleBySerial(ctx, *role.ParentSerial)
	if err != nil && !errutil.IsNotFoundError(err) {
		return err
	}

	if parentBySerial == nil {
		return errutil.NewBadRequestError("supervisor/parent is invalid")
	}
	return nil
}

func (u *Usecase) validatePermissionCodes(ctx context.Context, role *entities.Role, codes []string) ([]*entities.RolePermission, error) {
	var uniqueCodes []string
	mapCodes := map[string]bool{}
	for _, serial := range codes {
		if _, ok := mapCodes[serial]; !ok {
			uniqueCodes = append(uniqueCodes, serial)
			mapCodes[serial] = true
		}
	}

	mapPermission, err := u.permissionRepo.GetMapPermissionByCode(ctx, uniqueCodes)
	if err != nil {
		return nil, err
	}

	for _, code := range uniqueCodes {
		if _, ok := mapPermission[code]; !ok {
			return nil, errutil.NewBadRequestError(fmt.Sprintf("%s permission code is invalid", code))
		}
	}

	var rolePermissions []*entities.RolePermission
	for _, permission := range mapPermission {
		rolePermissions = append(rolePermissions, &entities.RolePermission{
			RoleSerial:     role.Serial,
			PermissionCode: permission.Code,
		})
	}

	return rolePermissions, nil
}

func (u *Usecase) appendRoleRelations(ctx context.Context, roles []*entities.Role) error {
	mapRoleBySerial := map[string]*entities.Role{}
	var roleSerials []string
	for _, role := range roles {
		roleSerials = append(roleSerials, role.Serial)
		mapRoleBySerial[role.Serial] = role
	}

	rolePermissions, err := u.roleRepo.GetRolePermissionsByRoleSerials(ctx, roleSerials)
	if err != nil {
		return err
	}

	var uniquePermissionCodes []string
	mapPermissionCode := map[string]bool{}
	for _, rolePermission := range rolePermissions {
		if ok := mapPermissionCode[rolePermission.PermissionCode]; !ok {
			uniquePermissionCodes = append(uniquePermissionCodes, rolePermission.PermissionCode)
			mapPermissionCode[rolePermission.PermissionCode] = true
		}

		//set relation role permissions
		role := mapRoleBySerial[rolePermission.RoleSerial]
		role.RolePermissions = append(role.RolePermissions, rolePermission)
	}

	mapPermissionByCode, err := u.permissionRepo.GetMapPermissionByCode(ctx, uniquePermissionCodes)
	if err != nil {
		return err
	}

	for _, role := range roles {
		//set relation permissions
		for _, rolePermission := range role.RolePermissions {
			if permission, ok := mapPermissionByCode[rolePermission.PermissionCode]; ok {
				role.Permissions = append(role.Permissions, permission)
			}
		}
	}
	return nil
}

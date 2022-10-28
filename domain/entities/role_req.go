package entities

import (
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"github.com/vodeacloud/hr-api/pkg/pagination"
)

type RolesResponse struct {
	Roles      []*Role                `json:"data"`
	Pagination *pagination.Pagination `json:"meta"`
}

type RolesResponseWithoutPagination struct {
	Roles []*Role `json:"data"`
}

type RoleResponse struct {
	Role *Role `json:"data"`
}

type GetRolesByQueryRequest struct {
	Page   int32  `form:"page"`
	Limit  int32  `form:"limit"`
	Sort   int32  `form:"sort"`
	Search string `form:"search"`
}

func (r *GetRolesByQueryRequest) Validate() error {
	if r.Page <= 0 {
		return errutil.NewBadRequestError("page must be start from 1")
	} else if r.Limit <= 0 {
		return errutil.NewBadRequestError("limit must be start from 1")
	}
	return nil
}

type GetRoleByIDRequest struct {
	ID int64 `uri:"id"`
}

func (r *GetRoleByIDRequest) Validate() error {
	if r.ID <= 0 {
		return errutil.NewBadRequestError("id is invalid")
	}
	return nil
}

type CreateRoleRequest struct {
	Name            string   `form:"name"`
	IsActive        bool     `form:"isActive"`
	ParentSerial    *string  `form:"parentSerial"`
	PermissionCodes []string `form:"permissionCodes"`
}

func (r *CreateRoleRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.Name == "" {
		fields["name"] = "name field is mandatory"
	}
	return errutil.BadRequestOrNil(fields)
}

type UpdateRoleRequest struct {
	ID              int64    `uri:"id"`
	Name            string   `form:"name"`
	IsActive        bool     `form:"isActive"`
	ParentSerial    *string  `form:"parentSerial"`
	PermissionCodes []string `form:"permissionCodes"`
}

func (r *UpdateRoleRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.ID <= 0 {
		fields["id"] = "id is invalid"
	}
	if r.Name == "" {
		fields["name"] = "name field is mandatory"
	}
	return errutil.BadRequestOrNil(fields)
}

type DeleteRoleRequest struct {
	ID int64 `uri:"id"`
}

func (r *DeleteRoleRequest) Validate() error {
	if r.ID <= 0 {
		return errutil.NewBadRequestError("id is invalid")
	}
	return nil
}

type DeleteRoleBulkRequest struct {
	IDs []int64 `form:"ids"`
}

func (r *DeleteRoleBulkRequest) Validate() error {
	if len(r.IDs) == 0 {
		return errutil.NewBadRequestError("ids is required")
	}
	return nil
}

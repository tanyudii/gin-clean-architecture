package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/usecases"
	"github.com/vodeacloud/hr-api/pkg/api"
	"net/http"
)

func NewRegisterRoleAPI(
	r *gin.RouterGroup,
	roleUc usecases.RoleUsecase,
) RoleAPI {
	h := &roleAPI{roleUc: roleUc}
	h.Register(r)
	return h
}

type RoleAPI interface {
	Register(r *gin.RouterGroup)
	GetRolesByQuery(c *gin.Context)
	GetRoleByID(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
	DeleteRoleBulk(c *gin.Context)
}

type roleAPI struct {
	roleUc usecases.RoleUsecase
}

func (h *roleAPI) Register(r *gin.RouterGroup) {
	r.GET("/v1/roles", h.GetRolesByQuery)
	r.GET("/v1/roles/:id", h.GetRoleByID)
	r.POST("/v1/roles", h.CreateRole)
	r.PUT("/v1/roles/:id", h.UpdateRole)
	r.DELETE("/v1/roles/:id", h.DeleteRole)
	r.POST("/v1/roles/delete-bulk", h.DeleteRoleBulk)
}

func (h *roleAPI) GetRolesByQuery(c *gin.Context) {
	var req entities.GetRolesByQueryRequest
	if err := c.ShouldBind(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.roleUc.GetRolesByQuery(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponseWithPagination(c, http.StatusOK, resp.Roles, resp.Pagination)
	return
}

func (h *roleAPI) GetRoleByID(c *gin.Context) {
	var req entities.GetRoleByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.roleUc.GetRoleByID(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponse(c, http.StatusOK, resp.Role)
	return
}

func (h *roleAPI) CreateRole(c *gin.Context) {
	var req entities.CreateRoleRequest
	if err := c.ShouldBind(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.roleUc.CreateRole(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponse(c, http.StatusCreated, resp.Role)
	return
}

func (h *roleAPI) UpdateRole(c *gin.Context) {
	var req entities.UpdateRoleRequest
	if err := c.ShouldBindUri(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.roleUc.UpdateRole(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponse(c, http.StatusOK, resp.Role)
	return
}

func (h *roleAPI) DeleteRole(c *gin.Context) {
	var req entities.DeleteRoleRequest
	if err := c.ShouldBindUri(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.roleUc.DeleteRole(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponse(c, http.StatusOK, resp.Role)
	return
}

func (h *roleAPI) DeleteRoleBulk(c *gin.Context) {
	var req entities.DeleteRoleBulkRequest
	if err := c.ShouldBind(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.roleUc.DeleteRoleBulk(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponse(c, http.StatusOK, resp.Roles)
	return
}

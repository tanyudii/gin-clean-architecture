package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
	"github.com/vodeacloud/hr-api/pkg/api"
	"net/http"
)

func NewRegisterPermissionAPI(
	r *gin.RouterGroup,
	permissionUc usecases.PermissionUsecase,
) PermissionAPI {
	h := &permissionAPI{permissionUc: permissionUc}
	h.Register(r)
	return h
}

type PermissionAPI interface {
	Register(r *gin.RouterGroup)
	GetAllPermissions(c *gin.Context)
}

type permissionAPI struct {
	permissionUc usecases.PermissionUsecase
}

func (h *permissionAPI) Register(r *gin.RouterGroup) {
	r.GET("/v1/permissions", h.GetAllPermissions)
}

func (h *permissionAPI) GetAllPermissions(c *gin.Context) {
	resp, err := h.permissionUc.GetAllPermissions(c)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponse(c, http.StatusOK, resp.Permissions)
	return
}

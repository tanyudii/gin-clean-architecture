package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

func NewRegisterGroupAPI(
	r *gin.RouterGroup,
	groupUc usecases.GroupUsecase,
) GroupAPI {
	h := &groupAPI{groupUc: groupUc}
	h.Register(r)
	return h
}

type GroupAPI interface {
	Register(r *gin.RouterGroup)
	GetGroupsByQuery(c *gin.Context)
	GetGroupByID(c *gin.Context)
	CreateGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
	DeleteGroupBulk(c *gin.Context)
}

type groupAPI struct {
	groupUc usecases.GroupUsecase
}

func (h *groupAPI) Register(r *gin.RouterGroup) {
	r.GET("/v1/groups", h.GetGroupsByQuery)
	r.GET("/v1/groups/:id", h.GetGroupByID)
	r.POST("/v1/groups", h.CreateGroup)
	r.PUT("/v1/groups/:id", h.UpdateGroup)
	r.DELETE("/v1/groups/:id", h.DeleteGroup)
	r.POST("/v1/groups/bulk-delete", h.DeleteGroupBulk)
}

func (h *groupAPI) GetGroupsByQuery(c *gin.Context) {

}

func (h *groupAPI) GetGroupByID(c *gin.Context) {

}

func (h *groupAPI) CreateGroup(c *gin.Context) {

}

func (h *groupAPI) UpdateGroup(c *gin.Context) {

}

func (h *groupAPI) DeleteGroup(c *gin.Context) {

}

func (h *groupAPI) DeleteGroupBulk(c *gin.Context) {

}

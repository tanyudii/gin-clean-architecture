package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

func NewRegisterUserAPI(
	r *gin.RouterGroup,
	userUc usecases.UserUsecase,
) UserAPI {
	h := &userAPI{userUc: userUc}
	h.Register(r)
	return h
}

type UserAPI interface {
	Register(r *gin.RouterGroup)
	GetUsersByQuery(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	DeleteUserBulk(c *gin.Context)
}

type userAPI struct {
	userUc usecases.UserUsecase
}

func (h *userAPI) Register(r *gin.RouterGroup) {
	r.GET("/v1/users", h.GetUsersByQuery)
	r.GET("/v1/users/:id", h.GetUserByID)
	r.POST("/v1/users", h.CreateUser)
	r.PUT("/v1/users/:id", h.UpdateUser)
	r.DELETE("/v1/users/:id", h.DeleteUser)
	r.POST("/v1/users/bulk-delete", h.DeleteUserBulk)
}

func (h *userAPI) GetUsersByQuery(c *gin.Context) {

}

func (h *userAPI) GetUserByID(c *gin.Context) {

}

func (h *userAPI) CreateUser(c *gin.Context) {

}

func (h *userAPI) UpdateUser(c *gin.Context) {

}

func (h *userAPI) DeleteUser(c *gin.Context) {

}

func (h *userAPI) DeleteUserBulk(c *gin.Context) {

}

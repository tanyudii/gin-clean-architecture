package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

func NewRegisterPasswordAPI(
	r *gin.RouterGroup,
	passwordUc usecases.PasswordUsecase,
) PasswordAPI {
	h := &passwordAPI{passwordUc: passwordUc}
	h.Register(r)
	return h
}

type PasswordAPI interface {
	Register(r *gin.RouterGroup)
	PasswordEmail(c *gin.Context)
	PasswordReset(c *gin.Context)
}

type passwordAPI struct {
	passwordUc usecases.PasswordUsecase
}

func (h *passwordAPI) Register(r *gin.RouterGroup) {
	r.POST("/v1/password/email", h.PasswordEmail)
	r.POST("/v1/password/reset", h.PasswordReset)
}

func (h *passwordAPI) PasswordEmail(c *gin.Context) {

}

func (h *passwordAPI) PasswordReset(c *gin.Context) {

}

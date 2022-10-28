package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/usecases"
	"github.com/vodeacloud/hr-api/pkg/api"
	"net/http"
)

func NewRegisterOAuthAPI(
	r *gin.RouterGroup,
	authUc usecases.OAuthUsecase,
) OAuthAPI {
	h := &oauthAPI{oauthUc: authUc}
	h.Register(r)
	return h
}

type OAuthAPI interface {
	Register(r *gin.RouterGroup)
	CreateToken(c *gin.Context)
}

type oauthAPI struct {
	oauthUc usecases.OAuthUsecase
}

func (h *oauthAPI) Register(r *gin.RouterGroup) {
	r.POST("/oauth/token", h.CreateToken)
}

func (h *oauthAPI) CreateToken(c *gin.Context) {
	var req entities.OAuthCreateTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	resp, err := h.oauthUc.CreateToken(c, &req)
	if err != nil {
		api.GinErrorResponse(c, err)
		return
	}

	api.GinResponseWithoutData(c, http.StatusOK, resp)
	return
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

func NewRegisterCompanyAPI(
	r *gin.RouterGroup,
	companyUc usecases.CompanyUsecase,
) CompanyAPI {
	h := &companyApi{companyUc: companyUc}
	h.Register(r)
	return h
}

type CompanyAPI interface {
	Register(r *gin.RouterGroup)
}

type companyApi struct {
	companyUc usecases.CompanyUsecase
}

func (h *companyApi) Register(r *gin.RouterGroup) {
}

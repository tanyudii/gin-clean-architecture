package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func NewRegisterHealthAPI(
	r *gin.RouterGroup,
) HealthAPI {
	h := &healthAPI{}
	h.Register(r)
	return h
}

type HealthAPI interface {
	Register(r *gin.RouterGroup)
	HealthCheck(c *gin.Context)
}

type healthAPI struct {
	db *gorm.DB
}

func (h *healthAPI) Register(r *gin.RouterGroup) {
	r.GET("_health", h.HealthCheck)
}

func (h *healthAPI) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

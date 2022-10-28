package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

func NewRegisterNotificationConfigAPI(
	r *gin.RouterGroup,
	notificationConfigUc usecases.NotificationConfigUsecase,
) NotificationConfigAPI {
	h := &notificationConfigAPI{notificationConfigUc: notificationConfigUc}
	h.Register(r)
	return h
}

type NotificationConfigAPI interface {
	Register(r *gin.RouterGroup)
	GetNotificationConfigs(c *gin.Context)
	UpsertNotificationConfigs(c *gin.Context)
}

type notificationConfigAPI struct {
	notificationConfigUc usecases.NotificationConfigUsecase
}

func (h *notificationConfigAPI) Register(r *gin.RouterGroup) {
	r.GET("/v1/notification-configs", h.GetNotificationConfigs)
	r.POST("/v1/notification-configs", h.UpsertNotificationConfigs)
}

func (h *notificationConfigAPI) GetNotificationConfigs(c *gin.Context) {

}

func (h *notificationConfigAPI) UpsertNotificationConfigs(c *gin.Context) {

}

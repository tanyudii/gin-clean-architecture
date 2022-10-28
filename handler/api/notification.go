package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

func NewRegisterNotificationAPI(
	r *gin.RouterGroup,
	notificationUc usecases.NotificationUsecase,
) NotificationAPI {
	h := &notificationAPI{notificationUc: notificationUc}
	h.Register(r)
	return h
}

type NotificationAPI interface {
	Register(r *gin.RouterGroup)
}

type notificationAPI struct {
	notificationUc usecases.NotificationUsecase
}

func (h *notificationAPI) Register(r *gin.RouterGroup) {

}

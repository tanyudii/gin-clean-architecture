package notification

import (
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	notificationRepo       repositories.NotificationRepository
	notificationConfigRepo repositories.NotificationConfigRepository
}

func NewUsecase(
	notificationRepo repositories.NotificationRepository,
	notificationConfigRepo repositories.NotificationConfigRepository,
) usecases.NotificationUsecase {
	return &Usecase{
		notificationRepo:       notificationRepo,
		notificationConfigRepo: notificationConfigRepo,
	}
}

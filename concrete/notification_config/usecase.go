package notification_config

import (
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
)

type Usecase struct {
	notificationConfigRepo repositories.NotificationConfigRepository
}

func NewUsecase(
	notificationConfigRepo repositories.NotificationConfigRepository,
) usecases.NotificationUsecase {
	return &Usecase{
		notificationConfigRepo: notificationConfigRepo,
	}
}

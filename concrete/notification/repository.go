package notification

import (
	"github.com/vodeacloud/hr-api/domain/repositories"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) repositories.NotificationRepository {
	return Repository{
		db: db,
	}
}

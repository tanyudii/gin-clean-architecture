package entities

import "time"

type Notification struct {
	ID         int64            `json:"id"`
	Serial     string           `json:"serial"`
	UserSerial string           `json:"userSerial"`
	Data       NotificationData `json:"data"`
	ReadAt     *time.Time       `json:"readAt" json:"readAt"`
	CreatedAt  *time.Time       `json:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time       `json:"updatedAt" json:"updatedAt"`
}

func (n Notification) TableName() string {
	return "notifications"
}

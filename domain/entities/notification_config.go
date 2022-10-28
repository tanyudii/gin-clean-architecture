package entities

type NotificationConfig struct {
	ID         int64
	UserSerial string
	Code       string
	Enabled    bool
}

func (n NotificationConfig) TableName() string {
	return "notification_configs"
}

package entities

import "time"

type User struct {
	ID         string      `json:"id"`
	Serial     string      `json:"serial"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	Phone      *string     `json:"phone"`
	AvatarURL  *string     `json:"avatarURL"`
	Type       UserType    `json:"type"`
	CreatedAt  *time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time  `json:"updatedAt"`
	CreatedBy  string      `json:"createdBy"`
	UpdatedBy  string      `json:"updatedBy"`
	UserDetail *UserDetail `json:"userDetail" gorm:"-"`
}

func (u User) TableName() string {
	return "users"
}

type UserDetail struct {
	ID            int64  `json:"id"`
	UserSerial    string `json:"userSerial"`
	CompanySerial string `json:"companySerial"`
}

func (u UserDetail) TableName() string {
	return "user_details"
}

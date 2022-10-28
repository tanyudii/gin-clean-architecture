package entities

import "time"

type Group struct {
	ID            int64        `json:"id"`
	Serial        string       `json:"serial"`
	CompanySerial string       `json:"-"`
	Name          string       `json:"name"`
	IsActive      bool         `json:"isActive"`
	ParentSerial  *string      `json:"parentSerial,omitempty"`
	CreatedAt     *time.Time   `json:"createdAt"`
	UpdatedAt     *time.Time   `json:"updatedAt"`
	CreatedBy     string       `json:"createdBy"`
	UpdatedBy     string       `json:"updatedBy"`
	Users         []*User      `json:"-" gorm:"-"`
	GroupUsers    []*GroupUser `json:"-" gorm:"-"`
}

func (g Group) TableName() string {
	return "groups"
}

type GroupUser struct {
	ID          int64  `json:"id"`
	GroupSerial string `json:"groupSerial"`
	UserSerial  string `json:"userSerial"`
}

func (g GroupUser) TableName() string {
	return "group_users"
}

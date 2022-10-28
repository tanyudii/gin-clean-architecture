package entities

import "time"

type Role struct {
	ID              int64             `json:"id"`
	CompanySerial   string            `json:"-"`
	Serial          string            `json:"serial"`
	Name            string            `json:"name"`
	IsActive        bool              `json:"isActive"`
	ParentSerial    *string           `json:"parentSerial,omitempty"`
	CreatedAt       *time.Time        `json:"createdAt"`
	UpdatedAt       *time.Time        `json:"updatedAt"`
	CreatedBy       string            `json:"createdBy"`
	UpdatedBy       string            `json:"updatedBy"`
	Permissions     []*Permission     `json:"-" gorm:"-"`
	RolePermissions []*RolePermission `json:"-" gorm:"-"`
}

func (r Role) TableName() string {
	return "roles"
}

type RolePermission struct {
	ID             int64  `json:"id"`
	RoleSerial     string `json:"roleSerial"`
	PermissionCode string `json:"permissionCode"`
}

func (r *RolePermission) TableName() string {
	return "role_permissions"
}

type RoleUser struct {
	ID         int64  `json:"id"`
	RoleSerial string `json:"roleSerial"`
	UserSerial string `json:"userSerial"`
}

func (r RoleUser) TableName() string {
	return "role_users"
}

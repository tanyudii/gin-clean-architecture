package entities

type Permission struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Transaction string `json:"transaction"`
}

func (p Permission) TableName() string {
	return "permissions"
}

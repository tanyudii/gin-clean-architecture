package entities

import "time"

type Company struct {
	ID        int64      `json:"id"`
	Serial    string     `json:"serial"`
	Name      string     `json:"name"`
	Initial   string     `json:"initial"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Fax       string     `json:"fax"`
	Address   string     `json:"address"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	ZipCode   string     `json:"zipCode"`
	IsActive  bool       `json:"isActive"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}

func (c Company) TableName() string {
	return "companies"
}

type CompanyDetail struct {
	ID            int64  `json:"id"`
	CompanySerial string `json:"companySerial"`
}

func (c CompanyDetail) TableName() string {
	return "company_details"
}

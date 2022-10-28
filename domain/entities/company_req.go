package entities

import "github.com/vodeacloud/hr-api/pkg/errutil"

type RegisterCompanyResponse struct {
	Password string
}

type RegisterCompanyRequest struct {
	Name          string `json:"name"`
	AdminEmail    string `json:"adminEmail"`
	AdminPassword string `json:"adminPassword"`
}

func (r *RegisterCompanyRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.Name == "" {
		fields["name"] = "name field is mandatory"
	}
	if r.AdminEmail == "" {
		fields["adminEmail"] = "admin email field is mandatory"
	}
	if r.AdminPassword == "" {
		fields["adminPassword"] = "admin password field is mandatory"
	}
	return errutil.BadRequestOrNil(fields)
}

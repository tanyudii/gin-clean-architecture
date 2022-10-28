package entities

import "github.com/vodeacloud/hr-api/pkg/errutil"

type UpsertNotificationConfigResponse struct {
	Config map[string]bool `json:"config"`
}

type UpsertNotificationConfigRequest struct {
	Config map[string]bool `form:"config"`
}

func (r *UpsertNotificationConfigRequest) Validate() error {
	fields := errutil.ErrorField{}
	if len(r.Config) == 0 {
		fields["config"] = "config field is required"
	}
	return errutil.BadRequestOrNil(fields)
}

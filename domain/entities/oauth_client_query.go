package entities

import "github.com/vodeacloud/hr-api/pkg/pagination"

type OAuthClientsQuery struct {
	Search     string
	Sort       int32
	Pagination *pagination.Pagination
}

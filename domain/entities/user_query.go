package entities

import "github.com/vodeacloud/hr-api/pkg/pagination"

type UsersQuery struct {
	Search     string
	Sort       int32
	Pagination *pagination.Pagination
}

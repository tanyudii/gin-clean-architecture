package pagination

import (
	"math"
)

const (
	DefaultPage  = 1
	DefaultLimit = 100
	MaximumLimit = 1000
)

type Pagination struct {
	Page      int32 `json:"page"`
	Limit     int32 `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int32 `json:"totalPage"`
}

func NewPagination(page int32, limit int32) *Pagination {
	p := &Pagination{Page: page, Limit: limit}
	p.Validate()
	return p
}

func (p *Pagination) Validate() *Pagination {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.Limit <= 0 || p.Limit > MaximumLimit {
		p.Limit = DefaultLimit
	}
	return p
}

func (p *Pagination) SetPagination() {
	if p.Total == 0 {
		return
	}
	p.TotalPage = int32(math.Ceil(float64(p.Total) / float64(p.Limit)))
}

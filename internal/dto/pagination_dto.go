package dto

type PaginationType string

const (
	PageBased   PaginationType = "page"
	CursorBased PaginationType = "cursor"
)

type PaginationQuery struct {
	Type       PaginationType `form:"type" validate:"omitempty,oneof=page cursor" default:"page" json:"type"`
	Page       int            `form:"page" validate:"gte=1" default:"1" json:"page"`
	PerPage    int            `form:"per_page" validate:"max=100" default:"10" json:"per_page"`
	LastSeenID string         `form:"last_seen_id" json:"last_seen_id"` // For cursor-based
	Asc        bool           `form:"asc" json:"asc"`                   // true=ASC, false=DESC
	SortField  string         `form:"sort_field" default:"_id" json:"sort_field"`
}

func (p *PaginationQuery) SetDefaults() {
	if p.Type == "" {
		p.Type = PageBased
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	if p.SortField == "" {
		p.SortField = "_id"
	}
}

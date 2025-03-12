package model

type Pagination struct {
	Total   int64 `json:"total"`
	Limit   int64 `json:"limit"`
	Page    int64 `json:"page"`
	HasMore bool  `json:"has_more"`
}

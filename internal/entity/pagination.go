package entity

type Pagination struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

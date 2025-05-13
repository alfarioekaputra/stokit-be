package model

type PaginatedResponse[T any] struct {
	Items []T   `json:"items"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
	First bool  `json:"first"`
	Last  bool  `json:"last"`
}

package repository

import (
	"stokit/internal/model"

	"net/http"

	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

func FetchAllWithFilter[T any, F any](
	db *gorm.DB,
	req *http.Request,
	filter *F,
	applyFilter func(*gorm.DB, *F) *gorm.DB,
) (*model.PaginatedResponse[T], error) {
	var items []T

	pg := paginate.New()
	query := db.Model(new(T))
	query = applyFilter(query, filter)

	pager := pg.With(query).Request(req).Response(&items)

	// Hitung halaman terakhir
	totalPages := (pager.Total + pager.Size - 1) / pager.Size
	isFirst := pager.Page == 1
	isLast := pager.Page >= totalPages

	return &model.PaginatedResponse[T]{
		Items: items,
		Page:  pager.Page,
		Size:  pager.Size,
		Total: pager.Total,
		First: isFirst,
		Last:  isLast,
	}, nil
}

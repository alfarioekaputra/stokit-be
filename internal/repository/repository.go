package repository

import (
	"net/http"
	"stokit/internal/model"

	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func FetchAllWithFilter[T any, F any](
	db *gorm.DB,
	req *http.Request,
	filter *F,
	applyFilter func(*gorm.DB, *F) *gorm.DB,
	applyPreload func(*gorm.DB) *gorm.DB,
) (*model.PaginatedResponse[T], error) {
	var items []T

	pg := paginate.New()
	query := db.Model(new(T))

	// Apply optional preloads
	if applyPreload != nil {
		query = applyPreload(query)
	}

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

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) CountById(db *gorm.DB, id any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error

	return total, err
}

func (r *Repository[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}

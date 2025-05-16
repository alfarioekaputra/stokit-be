package repository

import (
	"stokit/internal/entity"
	"stokit/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Repository[entity.Category]
	Log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{
		Log: log,
	}
}

func (r *CategoryRepository) GetTree(db *gorm.DB) ([]*entity.Category, error) {
	var categories []*entity.Category
	err := db.Find(&categories).Error

	return categories, err

}

func ApplyCategoryFilter(db *gorm.DB, filter *model.CategoryFilter) *gorm.DB {
	if filter.Name != nil {
		db = db.Where("name LIKE ?", "%"+*filter.Name+"%")
	}
	return db
}

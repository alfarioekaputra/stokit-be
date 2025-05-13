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

func ApplyCategoryFilter(db *gorm.DB, filter *model.CategoryFilter) *gorm.DB {
	if filter.Name != nil {
		db = db.Where("email = ?", *filter.Name)
	}
	if filter.ParentID != nil {
		db = db.Where("username = ?", *filter.ParentID)
	}
	return db
}

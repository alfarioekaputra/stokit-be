package repository

import (
	"stokit/internal/entity"
	"stokit/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	Repository[entity.Supplier]
	Log *logrus.Logger
}

func NewSupplierRepository(log *logrus.Logger) *SupplierRepository {
	return &SupplierRepository{
		Log: log,
	}
}

func ApplySupplierFilter(db *gorm.DB, filter *model.SupplierFilter) *gorm.DB {
	if filter.Name != nil {
		db = db.Where("name LIKE ?", "%"+*filter.Name+"%")
	}
	if filter.Address != nil {
		db = db.Where("address LIKE ?", "%"+*filter.Address+"%")
	}
	if filter.Phone != nil {
		db = db.Where("phone LIKE ?", "%"+*filter.Phone+"%")
	}
	return db
}

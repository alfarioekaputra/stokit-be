package repository

import (
	"stokit/internal/entity"

	"github.com/sirupsen/logrus"
)

type ProductRepository struct {
	Repository[entity.Product]
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

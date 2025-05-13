package usecase

import (
	"stokit/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUsecase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ProductRepository *repository.ProductRepository
}

func NewProductUsecase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, productRepository *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		ProductRepository: productRepository,
	}
}

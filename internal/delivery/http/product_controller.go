package http

import (
	"stokit/internal/usecase"

	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Log            *logrus.Logger
	ProductUsecase *usecase.ProductUsecase
}

func NewProductController(useCase *usecase.ProductUsecase, logger *logrus.Logger) *ProductController {
	return &ProductController{
		Log:            logger,
		ProductUsecase: useCase,
	}
}

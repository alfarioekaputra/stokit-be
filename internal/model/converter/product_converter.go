package converter

import (
	"stokit/internal/entity"
	"stokit/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Barcode:      product.Barcode,
		Brand:        product.Brand,
		ImageURL:     product.ImageURL,
		SKU:          product.SKU,
		CostPrice:    product.CostPrice,
		SellingPrice: product.SellingPrice,
		Stock:        product.Stock,
		Category:     product.Category.Name,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}
}

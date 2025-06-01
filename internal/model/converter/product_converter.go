package converter

import (
	"stokit/internal/entity"
	"stokit/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		ID:               product.ID,
		Name:             product.Name,
		Barcode:          product.Barcode,
		ImageURL:         product.ImageURL,
		SKU:              product.SKU,
		CostPrice:        product.CostPrice,
		SellingPrice:     product.SellingPrice,
		Stock:            product.Stock,
		LowStockTreshold: product.LowStockTreshold,
		Category:         product.Category.Name,
		CreatedAt:        product.CreatedAt,
		UpdatedAt:        product.UpdatedAt,
	}
}

package converter

import (
	"stokit/internal/entity"
	"stokit/internal/model"
)

func SupplierToResponse(supplier *entity.Supplier) *model.SupplierResponse {
	var products []model.ProductResponse
	for _, product := range supplier.Products {
		products = append(products, *ProductToResponse(&product))
	}
	return &model.SupplierResponse{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Address:   supplier.Address,
		Phone:     supplier.Phone,
		Email:     supplier.Email,
		Products:  products,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
}

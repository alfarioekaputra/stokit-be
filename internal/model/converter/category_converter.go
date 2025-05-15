package converter

import (
	"stokit/internal/entity"
	"stokit/internal/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		ParentID: *category.ParentID,
	}
}

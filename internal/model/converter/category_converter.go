package converter

import (
	"stokit/internal/entity"
	"stokit/internal/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	var parentID string
	var parentName string

	if category.ParentID != nil {
		parentID = *category.ParentID
	}

	if category.Parent != nil {
		parentName = category.Parent.Name
	}
	return &model.CategoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		ParentID: parentID,
		Parent:   parentName,
	}
}

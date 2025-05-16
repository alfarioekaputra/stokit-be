package model

type CategoryResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ParentID  string `json:"parent_id"`
	Parent    string `json:"parent"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type SearchCategoryRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"max=100"`
	Page int    `json:"page" validate:"min=1"`
	Size int    `json:"size" validate:"min=1,max=100"`
}

type CategoryFilter struct {
	Name     *string `query:"name"`
	ParentID *string `query:"parent_id"`
}

type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	ParentID string `json:"parent_id"`
}

type ViewCategoryRequest struct {
	ID string `json:"-" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID       string `json:"-" validate:"required"`
	Name     string `json:"name" validate:"required,max=100"`
	ParentID string `json:"parent_id"`
}

type DeleteCategoryRequest struct {
	ID string `json:"-" validate:"required"`
}

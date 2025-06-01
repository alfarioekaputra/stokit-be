package model

type SupplierResponse struct {
	ID        string            `json:"id,omitempty"`
	Name      string            `json:"name,omitempty"`
	Address   string            `json:"address,omitempty"`
	Phone     string            `json:"phone,omitempty"`
	Email     string            `json:"email,omitempty"`
	Products  []ProductResponse `json:"products,omitempty"`
	CreatedAt int64             `json:"created_at,omitempty"`
	UpdatedAt int64             `json:"updated_at,omitempty"`
}

type SearchSupplierRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"max=100"`
	Page int    `json:"page" validate:"min=1"`
	Size int    `json:"size" validate:"min=1,max=100"`
}

type SupplierFilter struct {
	Name    *string `query:"name"`
	Address *string `query:"address"`
	Phone   *string `query:"phone"`
	Email   *string `query:"email"`
}

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required,max=100"`
	Address string `json:"address" validate:"required,max=100"`
	Phone   string `json:"phone" validate:"required,max=100"`
	Email   string `json:"email" validate:"required,max=100"`
}

type ViewSupplierRequest struct {
	ID string `json:"-" validate:"required"`
}

type UpdateSupplierRequest struct {
	ID      string `json:"-" validate:"required"`
	Name    string `json:"name" validate:"required,max=100"`
	Address string `json:"address" validate:"required,max=100"`
	Phone   string `json:"phone" validate:"required,max=100"`
	Email   string `json:"email" validate:"required,max=100"`
}

type DeleteSupplierRequest struct {
	ID string `json:"-" validate:"required"`
}

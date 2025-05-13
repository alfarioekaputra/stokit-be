package model

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	ID       string `json:"-" validate:"required,max=100"`
	Password string `json:"password,omitempty" validate:"max=100"`
	Username string `json:"username,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type SearchUserRequest struct {
	ID       string `json:"id" validate:"required"`
	Username string `json:"name" validate:"max=100"`
	Email    string `json:"email" validate:"max=200"`
	Page     int    `json:"page" validate:"min=1"`
	Size     int    `json:"size" validate:"min=1,max=100"`
}

type UserFilter struct {
	Email    *string `query:"email"`
	Username *string `query:"username"`
}

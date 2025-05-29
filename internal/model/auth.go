package model

type Auth struct {
	// Login user id
	ID string
}

type (
	Token struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)

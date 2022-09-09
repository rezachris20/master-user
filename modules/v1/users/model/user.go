package model

type (
	RegisterNewUserRequest struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UpdateUserRequest struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}
)

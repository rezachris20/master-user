package dto

type UserDTO struct {
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

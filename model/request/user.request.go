package request

import "time"

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserUpdateRequest struct {
	Name      string    `json:"name" validate:"required"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdateEmailRequest struct {
	Email     string    `json:"email" validate:"required,email"`
	UpdatedAt time.Time `json:"updated_at"`
}
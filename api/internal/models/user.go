package models

import (
	"register/internal/database/db"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type UpdateUserResponse struct {
	Message     string  `json:"message"`
	UpdatedUser db.User `json:"updated_user"`
}

type RegisterResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (lr LoginRequest) Validate() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required, is.Email),
		validation.Field(&lr.Password, validation.Required),
	)
}

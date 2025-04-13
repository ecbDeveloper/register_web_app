package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Roles string

const (
	RoleAdmin Roles = "admin"
	RoleUser  Roles = "user"
)

type JwtCustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   Roles     `json:"role"`
	jwt.RegisteredClaims
}

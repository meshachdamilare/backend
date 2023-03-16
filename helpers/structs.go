package helpers

import "github.com/golang-jwt/jwt"

type InputCreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

type AuthTokenJwtClaim struct {
	Email string
	Name  string
	jwt.StandardClaims
}

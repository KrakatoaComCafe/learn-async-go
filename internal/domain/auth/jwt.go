package auth

import "time"

type UserClaims struct {
	UserID string
	Email  string
}

type TokenService interface {
	GenerateToken(claims UserClaims, expiresIn time.Duration) (string, error)
	ValidateToken(token string) (*UserClaims, error)
}

package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krakatoa/learn-async-go/internal/domain/auth"
)

type jwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) auth.TokenService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) GenerateToken(claims auth.UserClaims, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims.UserID,
		"email":   claims.Email,
		"exp":     time.Now().Add(expiresIn).Unix(),
	})

	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenStr string) (*auth.UserClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// validate
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return &auth.UserClaims{
		UserID: claims["user_id"].(string),
		Email:  claims["email"].(string),
	}, nil
}

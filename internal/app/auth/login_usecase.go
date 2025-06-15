package auth

import (
	"errors"
	"time"

	"github.com/krakatoa/learn-async-go/internal/domain/auth"
)

type loginUseCase struct {
	tokenService auth.TokenService
}

func NewLoginUseCase(ts auth.TokenService) auth.LoginUseCase {
	return &loginUseCase{
		tokenService: ts,
	}
}

func (uc *loginUseCase) Login(req auth.LoginRequest) (auth.LoginResponse, error) {
	if req.Email != "admin@email.com" || req.Password != "123456" {
		return auth.LoginResponse{}, errors.New("invalid credentials")
	}

	claims := auth.UserClaims{
		UserID: "1",
		Email:  req.Email,
	}

	token, err := uc.tokenService.GenerateToken(claims, time.Minute*15)
	if err != nil {
		return auth.LoginResponse{}, err
	}

	return auth.LoginResponse{Token: token}, nil
}

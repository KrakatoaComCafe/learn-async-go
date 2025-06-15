package auth

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

type LoginUseCase interface {
	Login(req LoginRequest) (LoginResponse, error)
}

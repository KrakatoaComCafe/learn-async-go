package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakatoa/learn-async-go/internal/domain/auth"
)

type AuthHandler struct {
	loginUseCase auth.LoginUseCase
}

func NewAuthHandler(luc auth.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase: luc,
	}
}

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/login", h.Login)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.loginUseCase.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, resp)
}

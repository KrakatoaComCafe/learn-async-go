package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krakatoa/learn-async-go/internal/domain/auth"
)

type AuthMiddleware struct {
	tokenService auth.TokenService
}

func NewAuthMiddleware(ts auth.TokenService) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: ts,
	}
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed token"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := m.tokenService.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

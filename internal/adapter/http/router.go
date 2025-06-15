package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakatoa/learn-async-go/internal/adapter/http/middleware"
)

func NewRouter(messageHandler *MessageHandler, authHandler *AuthHandler, authMiddlware *middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	messageHandler.RegisterRoutes(r, authMiddlware)
	authHandler.RegisterRoutes(r)

	return r
}

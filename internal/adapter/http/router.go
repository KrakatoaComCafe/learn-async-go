package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(messageHandler *MessageHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	messageHandler.RegisterRoutes(r)

	return r
}

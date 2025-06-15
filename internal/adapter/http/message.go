package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakatoa/learn-async-go/internal/adapter/http/middleware"
	"github.com/krakatoa/learn-async-go/internal/app"
	"github.com/krakatoa/learn-async-go/internal/domain/auth"
)

type MessageHandler struct {
	svc app.MessageService
}

func NewMessageHandler(svc app.MessageService) *MessageHandler {
	return &MessageHandler{
		svc: svc,
	}
}

func (h *MessageHandler) RegisterRoutes(r *gin.Engine, authMiddleware *middleware.AuthMiddleware) {
	protected := r.Group("/messages")
	protected.Use(authMiddleware.Handle())

	protected.POST("/", h.PostMessage)
	protected.GET("/", h.GetMessages)
}

type messageInput struct {
	Text string `json:"text" binding:"required"`
}

func (h *MessageHandler) PostMessage(c *gin.Context) {
	var input messageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text is obrigatory"})
		return
	}

	claims := c.MustGet("claims").(*auth.UserClaims)
	email := claims.Email

	log.Printf("[HTTP] User [%s] New message received %+v", email, input.Text)

	if err := h.svc.SendAndStoreMessage(input.Text); err != nil {
		log.Printf("[HTTP] Falha ao publicar mensagem '%s': %v", input.Text, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving message"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	messages := h.svc.GetMessages()
	c.JSON(http.StatusOK, messages)
}

package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakatoa/learn-async-go/internal/app"
)

type MessageHandler struct {
	svc app.MessageService
}

func NewMessageHandler(svc app.MessageService) *MessageHandler {
	return &MessageHandler{
		svc: svc,
	}
}

func (h *MessageHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/messages", h.PostMessage)
	r.GET("/messages", h.GetMessages)
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
	log.Printf("[HTTP] New message received %+v", input.Text)

	if err := h.svc.SendAndStoreMessage(input.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving message"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	messages := h.svc.GetMessages()
	c.JSON(http.StatusOK, messages)
}

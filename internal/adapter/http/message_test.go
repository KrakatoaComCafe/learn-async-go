package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	adapter "github.com/krakatoa/learn-async-go/internal/adapter/http"
	"github.com/krakatoa/learn-async-go/internal/adapter/http/middleware"
	"github.com/krakatoa/learn-async-go/internal/domain"
	"github.com/krakatoa/learn-async-go/internal/domain/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMessageService struct {
	mock.Mock
}

func (s *MockMessageService) SaveMessage(text string) error {
	args := s.Called(text)
	return args.Error(0)
}
func (s *MockMessageService) SendAndStoreMessage(text string) error {
	args := s.Called(text)
	return args.Error(0)
}
func (s *MockMessageService) GetMessages() []domain.Message {
	return []domain.Message(nil)
}

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateToken(claims auth.UserClaims, expiresIn time.Duration) (string, error) {
	args := m.Called(claims, expiresIn)
	return args.String(0), args.Error(1)
}
func (m *MockTokenService) ValidateToken(token string) (*auth.UserClaims, error) {
	args := m.Called(token)
	return args.Get(0).(*auth.UserClaims), args.Error(1)
}

func TestMessageHandler_HandleMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// mocks
	mockService := new(MockMessageService)
	mockTokenService := new(MockTokenService)
	// handlers & middlewares
	handler := adapter.NewMessageHandler(mockService)
	authMiddlware := middleware.NewAuthMiddleware(mockTokenService)

	// mock token
	mockClaims := &auth.UserClaims{
		UserID: "test-user",
		Email:  "test-user@email.com",
	}
	mockTokenService.On("ValidateToken", "mock-token").Return(mockClaims, nil)

	router := gin.Default()
	handler.RegisterRoutes(router, authMiddlware)

	mockService.On("SendAndStoreMessage", "hello").Return(nil)

	payload := map[string]string{"text": "hello"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/messages/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer mock-token")
	// execute
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertCalled(t, "SendAndStoreMessage", "hello")
	mockTokenService.AssertCalled(t, "ValidateToken", "mock-token")
}

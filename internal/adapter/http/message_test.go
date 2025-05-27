package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	adapter "github.com/krakatoa/learn-async-go/internal/adapter/http"
	"github.com/krakatoa/learn-async-go/internal/domain"
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

func TestMessageHandler_HandleMessage(t *testing.T) {
	mockService := new(MockMessageService)
	handler := adapter.NewMessageHandler(mockService)
	router := adapter.NewRouter(handler)

	mockService.On("SendAndStoreMessage", "hello").Return(nil)

	payload := map[string]string{"text": "hello"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertCalled(t, "SendAndStoreMessage", "hello")
}

package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/armzerpa/travel-itinerary-generator/internal/handlers"
	"github.com/armzerpa/travel-itinerary-generator/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGenerateHandler(t *testing.T) {
	// Create a new Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the GenerateHandler
	handlers.GenerateHandler(c)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"GET request received"}`, w.Body.String())
}

// needs to be fixed
func TestPostHandler(t *testing.T) {
	// Create a request with a valid payload
	request := models.Request{
		Prompt:        "Write me an itinerary for 4 days in Rome",
		City:          "Rome",
		NumberOfLines: 4,
		Type:          "text",
	}
	requestBody, _ := json.Marshal(request)

	// Create a new Gin context with the request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/generate", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Mock the GeminiItineraryGenerator service
	//services.GeminiItineraryGenerator := &MockGeminiItineraryGenerator{}

	// Call the PostHandler
	handlers.PostHandler(c)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Itinerary for 4 days in Rome"}`, w.Body.String())
}

// MockGeminiItineraryGenerator is a mock implementation of the GeminiItineraryGenerator service
type MockGeminiItineraryGenerator struct{}

func (m *MockGeminiItineraryGenerator) GenerateItinerary(prompt, city string, numberOfLines int, typeArg string) (string, error) {
	return "Itinerary for 4 days in Rome", nil
}

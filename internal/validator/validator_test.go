package validator_test

import (
	"errors"
	"testing"

	"github.com/armzerpa/travel-itinerary-generator/internal/validator"
	"github.com/armzerpa/travel-itinerary-generator/models"
)

func TestValidateRequest(t *testing.T) {
	// Test case 1: Valid request
	req := &models.Request{
		// Set valid request fields here
		Prompt:        "prompt",
		City:          "city",
		NumberOfLines: 1,
		Type:          "none",
	}
	err := validator.ValidateRequest(req)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case 2: Invalid request
	req = &models.Request{
		// Set invalid request fields here
		Prompt: "",
	}
	expectedErr := errors.New("Prompt is required")
	err = validator.ValidateRequest(req)
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error message: %s, but got: %s", expectedErr.Error(), err.Error())
	}

	// Test case 3: Empty request
	req = &models.Request{
		// Set invalid request fields here
		Prompt: "prompt",
		City:   "",
	}
	expectedErr = errors.New("City is required")
	err = validator.ValidateRequest(req)
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error message: %s, but got: %s", expectedErr.Error(), err.Error())
	}

}

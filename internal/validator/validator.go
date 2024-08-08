package validator

import (
	"errors"

	"github.com/armzerpa/travel-itinerary-generator/models"
	"github.com/gin-gonic/gin"
)

func ValidateRequest(req *models.Request) error {
	if req.Prompt == "" {
		return gin.Error{Err: errors.New("Prompt is required"), Type: gin.ErrorTypeBind, Meta: nil}
	}

	if req.NumberOfLines < 1 || req.NumberOfLines > 100 {
		return gin.Error{Err: errors.New("NumberOfLines must be between 1 and 100"), Type: gin.ErrorTypeBind, Meta: nil}
	}

	if req.Type != "none" && req.Type != "outdoor" && req.Type != "indoor" {
		return gin.Error{Err: errors.New("type must be either 'none' or 'outdoor' or 'indoor'"), Type: gin.ErrorTypeBind, Meta: nil}
	}

	return nil
}

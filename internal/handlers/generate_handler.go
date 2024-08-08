package handlers

import (
	"net/http"

	"github.com/armzerpa/travel-itinerary-generator/internal/validator"
	"github.com/armzerpa/travel-itinerary-generator/models"
	"github.com/gin-gonic/gin"
)

func GenerateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET request received",
	})
}

func PostHandler(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate the request
	if err := validator.ValidateRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Process the request
	// ...

	c.JSON(http.StatusOK, gin.H{
		"message":       "POST request received",
		"prompt":        req.Prompt,
		"NumberOfLines": req.NumberOfLines,
		"type":          req.Type,
	})
}

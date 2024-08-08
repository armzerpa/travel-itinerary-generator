package handlers

import (
	"fmt"
	"net/http"

	"github.com/armzerpa/travel-itinerary-generator/internal/validator"
	"github.com/armzerpa/travel-itinerary-generator/models"
	"github.com/armzerpa/travel-itinerary-generator/services"
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

	generator := &services.MockItineraryGenerator{}

	// Call the generateItinerary method
	itinerary, err := generator.GenerateItinerary(req.Prompt, req.City, req.NumberOfLines, req.Type)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": itinerary,
	})
}

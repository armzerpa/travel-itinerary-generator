package main

import (
	"net/http"

	"github.com/armzerpa/travel-itinerary-generator/authentication"
	"github.com/gin-gonic/gin"
)

// Request struct
type Request struct {
	Prompt        string `json:"prompt"`
	NumberOfLines int    `json:"NumberOfLines"`
	Type          string `json:"type"`
}

func main() {
	r := gin.Default()
	r.Use(authentication.ApiKeyMiddleware)

	// GET route
	r.GET("/generate", generateHandler)

	// POST route
	r.POST("/generate", postHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func generateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET request received",
	})
}

func postHandler(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
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

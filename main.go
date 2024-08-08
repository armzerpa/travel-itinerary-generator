package main

import (
	"github.com/armzerpa/travel-itinerary-generator/internal/auth"
	"github.com/armzerpa/travel-itinerary-generator/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(auth.ApiKeyMiddleware)

	// GET route
	r.GET("/generate", handlers.GenerateHandler)

	// POST route
	r.POST("/generate", handlers.PostHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}

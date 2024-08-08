package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/armzerpa/travel-itinerary-generator/models"
)

const postUrl = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key="

type GeminiItineraryGenerator struct {
}

// generateItinerary is the implementation of the ItineraryGenerator interface for the Gemini model.
func (g *GeminiItineraryGenerator) GenerateItinerary(prompt string, city string, numberOfLines int, typeArg string) (string, error) {

	requestPayload := models.GeminiBody{
		Contents: []models.Content{
			{
				Parts: []models.Part{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	// Convert the request payload to JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY environment variable not set")
		return "", fmt.Errorf("API_KEY environment variable not set")
	}

	complete_url := fmt.Sprintf("%s%s", postUrl, apiKey)

	// Create the HTTP request
	req, err := http.NewRequest("POST", complete_url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	itinerary := string(respBody)
	return itinerary, nil
}

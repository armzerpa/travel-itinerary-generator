package services

import "fmt"

type MockItineraryGenerator struct {
}

// generateItinerary is the implementation of the ItineraryGenerator interface.
func (g *MockItineraryGenerator) GenerateItinerary(prompt string, city string, numberOfLines int, typeArg string) (string, error) {
	// Generate the itinerary based on the input parameters
	itinerary := fmt.Sprintf("Itinerary for '%s' (%d lines, type: %s)", city, numberOfLines, typeArg)
	return itinerary, nil
}

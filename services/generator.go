package services

type ItineraryGenerator interface {
	GenerateItinerary(prompt string, city string, numberOfLines int, typeArg string) (string, error)
}

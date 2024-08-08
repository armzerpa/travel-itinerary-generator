package models

// Request struct
type Request struct {
	Prompt        string `json:"prompt"`
	City          string `json:"city"`
	NumberOfLines int    `json:"NumberOfLines"`
	Type          string `json:"type"`
}

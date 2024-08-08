package models

// Request struct
type Request struct {
	Prompt        string `json:"prompt"`
	NumberOfLines int    `json:"NumberOfLines"`
	Type          string `json:"type"`
}

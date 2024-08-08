package models

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeminiBody struct {
	Contents []Content `json:"contents"`
}

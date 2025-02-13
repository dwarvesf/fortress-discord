package model

type AIResponse struct {
	Input    string
	Response string
}

type N8NEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type N8NEmbedFooter struct {
	Text string `json:"text"`
}

type N8NEmbedResponse struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Color       string          `json:"color"`
	Fields      []N8NEmbedField `json:"fields"`
	Footer      N8NEmbedFooter  `json:"footer"`
}

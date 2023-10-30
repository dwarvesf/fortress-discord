package model

// AdapterTrend is a struct response from adapter, before process to in-app model
type AdapterTrend struct {
	Data []*Repo `json:"data"`
}

// Trend is in-app model, after process from adapters
type Repo struct {
	Name                string `json:"name"`
	SpokenLanguage      string `json:"spoken_lang"`
	ProgrammingLanguage string `json:"program_lang"`
	StarCount           uint16 `json:"star_count"`
	ForkCount           uint16 `json:"fork_count"`
	StarGained          uint16 `json:"star_gained"`
	Description         string `json:"description"`
	DateRange           string `json:"date_range"`
	URL                 string `json:"url"`
	// CreatedAt           *time.Time `json:"created_at"`
}

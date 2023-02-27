package model

// ChangelogDigest is a struct response from adapter, before process to in-app model
type ChangelogDigest struct {
	Data    []*Changelog `json:"data"`
	Message string       `json:"message"`
}

// Changelog is a model for changelog
type Changelog struct {
	RowID        string `json:"row_id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	ChangelogURL string `json:"changelog_url"`
}

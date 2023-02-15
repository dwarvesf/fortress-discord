package model

type AdapterIssue struct {
	Data    []*Issue `json:"data"`
	Message string   `json:"message"`
}

type Issue struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	RootCause string `json:"root_cause"`
	Scope     string `json:"scope"`
}

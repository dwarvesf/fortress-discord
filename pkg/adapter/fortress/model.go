package fortress

type ErrorMessage struct {
	Message string `json:"message"`
}

type EmployeeSearch struct {
	DiscordID string
	Email     string
	Key       string
}

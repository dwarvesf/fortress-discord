package fortress

import "github.com/dwarvesf/fortress-discord/pkg/model"

type ErrorMessage struct {
	Message string `json:"message"`
}

type EmployeeSearch struct {
	DiscordID string
	Email     string
	Key       string
}

type SetSpeakerResponse struct {
	Data model.Event `json:"data"`
}

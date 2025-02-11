package ai

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type AI struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) AICommander {
	return &AI{
		L:    l,
		svc:  svc,
		view: view,
	}
}

func (a *AI) ProcessAI(message *model.DiscordMessage) error {
	input := strings.TrimSpace(strings.TrimPrefix(message.RawContent, "?ai"))

	if input == "" {
		return a.view.Error().Raise(message, "Please provide some text to process.")
	}

	// Process the text using the AI service
	response, err := a.svc.AI().ProcessTextWithN8N(input)
	if err != nil {
		a.L.Error(err, "failed to process AI text")
		return err
	}

	// Send the AI response back to the user
	return a.view.AI().SendResponse(message, response)
}

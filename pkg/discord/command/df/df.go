package df

import (
	"fmt"
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type DF struct {
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
	cfg  *config.Config
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, cfg *config.Config) DFCommander {
	return &DF{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

func (a *DF) ProcessWithN8N(message *model.DiscordMessage) error {
	input := strings.TrimSpace(strings.TrimPrefix(message.RawContent, "?df"))

	if input == "" {
		return a.view.Error().Raise(message, "Please provide some text to process.")
	}

	userSortRoles, err := a.svc.Profile().GetDiscordRoles(a.cfg.Discord.ID.DwarvesGuild, message.Author.ID)
	if err != nil {
		a.L.Error(err, "failed to get sort roles for author ID")
		return a.view.Error().Raise(message, fmt.Sprintf("Error: %v", err.Error()))
	}

	highestRole := ""
	if len(userSortRoles) > 0 {
		highestRole = userSortRoles[0]
	}

	// Process the text using the AI service
	response, err := a.svc.AI().ProcessTextWithN8N(input, message.Author.ID, message.Author.Username, highestRole)
	if err != nil {
		a.L.Error(err, "failed to process AI text with N8N")
		return a.view.Error().Raise(message, fmt.Sprintf("Error: %v", err.Error()))
	}

	// Send the AI response back to the user
	return a.view.DF().SendResponse(message, response)
}

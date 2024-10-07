package profile

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Profile struct {
	adapter adapter.IAdapter
	l       logger.Logger
	ses     *discordgo.Session
}

func New(adapter adapter.IAdapter, l logger.Logger, ses *discordgo.Session) Service {
	return &Profile{
		adapter: adapter,
		l:       l,
		ses:     ses,
	}
}

type Service interface {
	GetEmployeeList(in EmployeeSearch) (rs []model.Employee, err error)
	GetDiscordRoles(guildID string, userID string) (rs []string, err error)
}

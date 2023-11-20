package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command"
	"github.com/dwarvesf/fortress-discord/pkg/discord/history"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type Discord struct {
	Session *discordgo.Session

	Cfg *config.Config
	L   logger.Logger

	MessageHistory *history.MsgHistory

	Command *command.Command
}

func New(ses *discordgo.Session, cfg *config.Config, l logger.Logger, svc service.Servicer, view view.Viewer, msgHistory *history.MsgHistory) *Discord {
	return &Discord{
		Session:        ses,
		Cfg:            cfg,
		L:              l,
		MessageHistory: msgHistory,
		Command:        command.New(cfg, l, svc, view, msgHistory),
	}
}

func (d *Discord) ListenAndServe() (*discordgo.Session, error) {
	d.L.Info("open discord session")

	// register handlers
	d.Session.AddHandler(func(s *discordgo.Session, curMsg *discordgo.MessageCreate) {
		d.onMessageCreate(s, curMsg)
		d.MessageHistory.AddMsg(curMsg.Message)
	})

	// register interaction, right now use for sending updates
	d.Session.AddHandler(d.onInteractionCreate)
	d.Session.AddHandler(d.onReactionCreate)
	d.Session.AddHandler(d.onReactionRemove)
	d.Session.AddHandler(d.onAllReactionsRemove)

	// intents to receive message
	d.Session.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMessageReactions | discordgo.IntentsDirectMessages | discordgo.IntentsDirectMessageReactions

	err := d.Session.Open()
	if err != nil {
		d.L.Error(err, "failed to open discord session")
		return nil, err
	}
	d.L.Info("discord session opened")

	return d.Session, nil
}

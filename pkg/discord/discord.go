package discord

import (
	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
)

type Discord struct {
	Session *discordgo.Session

	Cfg *config.Config
	L   logger.Logger

	Command *command.Command
}

func New(cfg *config.Config, l logger.Logger) *Discord {
	discord, err := discordgo.New("Bot " + cfg.Discord.SecretToken)
	if err != nil {
		l.Fatal(err, "failed to create discord session")
	}
	return &Discord{
		Session: discord,
		Cfg:     cfg,
		L:       l,
		Command: command.New(l),
	}
}

func (d *Discord) ListenAndServe() (*discordgo.Session, error) {
	d.L.Info("open discord session")

	// register handlers
	d.Session.AddHandler(d.onMessageCreate)

	// intents to receive message
	d.Session.Identify.Intents = discordgo.IntentsGuildMessages

	err := d.Session.Open()
	if err != nil {
		d.L.Error(err, "failed to open discord session")
		return nil, err
	}
	d.L.Info("discord session opened")

	return d.Session, nil
}

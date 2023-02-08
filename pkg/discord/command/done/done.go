package done

import (
	"strings"

	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Done struct {
	cfg  *config.Config
	L    logger.Logger
	svc  service.Servicer
	view view.Viewer
}

func New(cfg *config.Config, l logger.Logger, svc service.Servicer, view view.Viewer) DoneCommander {
	return &Done{
		L:    l,
		svc:  svc,
		view: view,
		cfg:  cfg,
	}
}

func (d *Done) Done(message *model.DiscordMessage) error {
	// TODO: remove hard code
	icy := "0.3"
	icyF := 0.3

	// we check the content
	if len(message.ContentArgs) == 1 {
		d.L.Info("missing done content")
		return d.view.Done().MissingContent(message)
	}
	msg := strings.Join(message.ContentArgs[1:], " ")

	// first, we repost to a channel
	err := d.view.Done().Repost(message, msg, d.cfg.Discord.ID.RepostDoneChannel, icy)
	if err != nil {
		d.L.Error(err, "can't repost message")
		return err
	}

	// we send tip
	tip := &model.Tip{
		Recipients:   []string{message.Author.ID},
		Sender:       d.cfg.Discord.ID.FortressBot,
		Amount:       icyF,
		GuildId:      d.cfg.Discord.ID.DwarvesGuild,
		Token:        "icy",
		TransferType: "tip",
		Platform:     "discord",
		ChannelId:    d.cfg.Discord.ID.RepostDoneChannel,
	}

	err = d.svc.Treasury().SendTip(tip)
	if err != nil {
		d.L.Error(err, "can't send tip")
		return d.view.Done().CantSendReward(message)
	}

	return nil
}

package changelog

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/history"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Changelog struct {
	L          logger.Logger
	svc        service.Servicer
	view       view.Viewer
	msgHistory *history.MsgHistory
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer, msgHistory *history.MsgHistory) ChangelogCommander {
	return &Changelog{
		L:          l,
		svc:        svc,
		view:       view,
		msgHistory: msgHistory,
	}
}

func (t *Changelog) Send(message *model.DiscordMessage) error {
	// 1. get data from service
	data, err := t.svc.Changelog().GetListChangelogs()
	if err != nil {
		t.L.Error(err, "can't get list of Changelog")
		return err
	}

	// 2. render
	err = t.view.Changelog().Changelog(message, data)
	if err != nil {
		t.L.Error(err, "can't render Changelog")
		return err
	}

	// 3. handle history
	go t.msgHistory.NewFilter("send-changelog", func(receiveMsg *discordgo.Message) bool {
		// only accept message from the same user
		if receiveMsg.Author.ID != message.Author.ID {
			return false
		}

		// matches strings that contain one or more numbers separated by commas
		matchExpression := regexp.MustCompile(`^\d+(,\s*\d+)*$`)
		if matchExpression.MatchString(strings.ToLower(receiveMsg.Content)) {
			return true
		}
		return false
	}, 1, 30*time.Second, false,
		func(collectMsg []*discordgo.Message) {
			if len(collectMsg) == 0 {
				return
			}

			// 4. handle user input
			msg := collectMsg[0]
			inputs := []int{}
			for _, i := range strings.Split(msg.Content, ",") {
				// convert string to int
				n, err := strconv.Atoi(i)
				if err != nil {
					t.L.Error(err, "can't convert string to int")
					return
				}
				inputs = append(inputs, n)
			}

			// 5. send changelog one by one
			for _, i := range inputs {
				if i > len(data) {
					continue
				}
				c := data[i]

				if err := t.svc.Changelog().SendChangelog(c); err != nil {
					t.L.Error(err, "can't send Changelog")
					if err := t.view.Changelog().ChangelogSendFailed(message, c); err != nil {
						t.L.Error(err, "can't render Changelog send success")
						return
					}
					return
				}

				if err := t.view.Changelog().ChangelogSendSuccess(message, c); err != nil {
					t.L.Error(err, "can't render Changelog send success")
					return
				}
			}
		})

	return nil
}

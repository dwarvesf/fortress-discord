package command

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/adopt"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/assess"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/digest"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/help"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/hold"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/memo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/milestone"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/new"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/trial"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/updates"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Command struct {
	L    logger.Logger
	Cmds map[string]base.TextCommander
	View view.Viewer
}

func New(l logger.Logger, svc service.Servicer, view view.Viewer) *Command {
	cmd := &Command{
		Cmds: make(map[string]base.TextCommander),
		L:    l,
		View: view,
	}

	// register all commands here
	cmd.Add([]base.TextCommander{
		earn.New(l, svc, view),
		help.New(l, svc, view),
		trial.New(l, svc, view),
		assess.New(l, svc, view),
		adopt.New(l, svc, view),
		hold.New(l, svc, view),
		new.New(l, svc, view),
		hiring.New(l, svc, view),
		event.New(l, svc, view),
		staff.New(l, svc, view),
		milestone.New(l, svc, view),
		digest.New(l, svc, view),
		updates.New(l, svc, view),
		memo.New(l, svc, view),
	})

	return cmd
}

func (c *Command) Add(cmds []base.TextCommander) error {
	// we register all commands to a map
	for i := range cmds {
		c.L.Infof("Registered command %s", cmds[i].Name())
		// support multiple prefixes for a command
		prefixes := cmds[i].Prefix()
		for _, p := range prefixes {
			c.Cmds[p] = cmds[i]
		}
	}
	return nil
}

func (c *Command) Execute(m *model.DiscordMessage) error {
	l := c.L.Fields(logger.Fields{
		"message": m,
	})

	// execute base on the prefix command
	cmd, ok := c.Cmds[m.ContentArgs[0]]
	if !ok {
		l.Info("command not found")
		return c.View.Error().CommandNotFound(m)
	}

	l.Field("cmd", cmd.Name()).Debug("execute command")
	err := cmd.Execute(m)
	if err != nil {
		// add custom generic todo here (something went wrong, permission, ..)
		return err
	}

	return nil
}

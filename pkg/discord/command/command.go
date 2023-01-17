package command

import (
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/earn"
	"github.com/dwarvesf/fortress-discord/pkg/errs"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Command struct {
	L    logger.Logger
	Cmds map[string]base.TextCommander
}

func New(l logger.Logger) *Command {
	cmd := &Command{
		Cmds: make(map[string]base.TextCommander),
		L:    l,
	}

	// register all commands here
	cmd.Add([]base.TextCommander{
		earn.New(l),
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
		return errs.ErrInvalidCommand
	}

	l.Field("cmd", cmd.Name()).Debug("execute command")
	return cmd.Execute(m)
}

package command

import (
	"github.com/dwarvesf/fortress-discord/pkg/config"
	"github.com/dwarvesf/fortress-discord/pkg/discord/base"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/adopt"
	salary "github.com/dwarvesf/fortress-discord/pkg/discord/command/advance"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/assess"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/brainery"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/changelog"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/deliverymetrics"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/digest"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/done"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/earn"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/event"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/help"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/hiring"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/hold"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/icy"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/index"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/issue"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/memo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/milestone"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/mma"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/new"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/news"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/profile"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/radar"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/staff"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/sum"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/topic"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/trend"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/trial"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/updates"
	"github.com/dwarvesf/fortress-discord/pkg/discord/command/withdraw"
	"github.com/dwarvesf/fortress-discord/pkg/discord/history"
	"github.com/dwarvesf/fortress-discord/pkg/discord/service"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Command struct {
	L    logger.Logger
	Cmds map[string]base.TextCommander
	View view.Viewer
	S    service.Servicer
}

func New(cfg *config.Config, l logger.Logger, svc service.Servicer, view view.Viewer, msgHistory *history.MsgHistory) *Command {
	cmd := &Command{
		Cmds: make(map[string]base.TextCommander),
		L:    l,
		View: view,
		S:    svc,
	}

	// register all commands here
	cmd.Add([]base.TextCommander{
		adopt.New(l, svc, view),
		assess.New(l, svc, view),
		brainery.New(l, svc, view, cfg),
		changelog.New(l, svc, view, msgHistory),
		deliverymetrics.New(l, svc, view, cfg),
		digest.New(l, svc, view),
		done.New(cfg, l, svc, view),
		earn.New(l, svc, view),
		event.New(l, svc, view),
		help.New(l, svc, view),
		hiring.New(l, svc, view),
		hold.New(l, svc, view),
		icy.New(l, svc, view, cfg),
		index.New(l, svc, view),
		issue.New(l, svc, view),
		memo.New(l, svc, view, cfg),
		milestone.New(l, svc, view),
		new.New(l, svc, view),
		radar.New(l, svc, view),
		staff.New(l, svc, view),
		sum.New(l, svc, view),
		trial.New(l, svc, view),
		updates.New(l, svc, view),
		profile.New(l, svc, view, cfg),
		mma.New(l, svc, view, cfg),
		trend.New(l, svc, view),
		salary.New(l, svc, view),
		withdraw.New(l, svc, view),
		news.New(l, svc, view, cfg),
		topic.New(l, svc, view, cfg),
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

	// run a permission check
	canExec, required := cmd.PermissionCheck(m)
	if !canExec {
		return c.View.Error().NotHavePermission(m, required)
	}

	l.Field("cmd", cmd.Name()).Debug("execute command")
	err := cmd.Execute(m)
	if err != nil {
		// add custom generic todo here (something went wrong, permission, ..)
		return err
	}

	return nil
}

package earn

import (
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Earn struct {
	L logger.Logger
}

func New(l logger.Logger) IEarn {
	return &Earn{
		L: l,
	}
}

func (e *Earn) List(message *model.DiscordMessage) error {
	l := e.L.Fields(logger.Fields{"message": message, "command": "list"})
	l.Info("executing command")
	defer l.Info("executed command")

	return nil
}

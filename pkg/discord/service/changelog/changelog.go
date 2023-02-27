package changelog

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Changelog struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) ChangelogServicer {
	return &Changelog{
		adapter: adapter,
		l:       l,
	}
}

func (c *Changelog) GetListChangelogs() ([]*model.Changelog, error) {
	// get response from fortress
	adapterChangelog, err := c.adapter.Fortress().GetChangelogs()
	if err != nil {
		c.l.Error(err, "can't get open changelog from fortress")
		return nil, err
	}

	// normalized into in-app model
	changelog := adapterChangelog.Data

	return changelog, nil
}

// SendChangelog implements ChangelogServicer
func (c *Changelog) SendChangelog(data *model.Changelog) error {
	return c.adapter.Fortress().SendChangelog(data)
}

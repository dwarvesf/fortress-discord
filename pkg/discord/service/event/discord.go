package event

import "github.com/dwarvesf/fortress-discord/pkg/model"

func (e *Event) CreateGuildScheduledEvent(ev *model.DiscordEvent) error {
	e.l.Field("name", ev.Name).Info("create a scheduled event on discord")
	// get response from fortress
	err := e.adapter.Fortress().CreateGuildScheduledEvent(ev)
	if err != nil {
		e.l.Error(err, "can't create a scheduled event on discord")
		return err
	}

	return nil
}

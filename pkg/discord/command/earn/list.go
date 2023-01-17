package earn

import "github.com/dwarvesf/fortress-discord/pkg/model"

func (e *Earn) List(message *model.DiscordMessage) error {
	// TODO: implement

	// 1. get data from service
	data, err := e.svc.Earn().GetActiveList()
	if err != nil {
		e.L.Error(err, "can't get list of active earn")
		return err
	}

	// 2. render
	err = e.view.Earn().List(data)
	if err != nil {
		return err
	}

	return nil
}

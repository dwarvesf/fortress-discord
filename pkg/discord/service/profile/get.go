package profile

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Profile) Get(id string) (*model.Employee, error) {
	rs, err := e.adapter.Fortress().GetEmployeeByDiscordID(id)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

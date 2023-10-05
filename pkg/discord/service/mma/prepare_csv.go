package mma

import (
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *MMA) GetEmployeeMMAScore() ([]model.EmployeeMMAScore, error) {
	employeeMMAScores, err := e.adapter.Fortress().GetEmployeesWithMMAScore()
	if err != nil {
		return nil, err
	}

	return employeeMMAScores, nil
}

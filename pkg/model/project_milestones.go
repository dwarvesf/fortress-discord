package model

import "time"

type AdapterProjectMilestone struct {
	Data    []*ProjectMilestone `json:"data"`
	Message string              `json:"message"`
}

type Milestone struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	StartDate     *time.Time  `json:"start_date"`
	EndDate       *time.Time  `json:"end_date"`
	SubMilestones []Milestone `json:"sub_milestones"`
}

type ProjectMilestone struct {
	Name       string      `json:"name"`
	Milestones []Milestone `json:"milestones"`
}

package profile

import (
	"sort"

	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EmployeeSearch struct {
	DiscordID string
	Email     string
	Key       string
}

func (e *Profile) GetEmployeeList(in EmployeeSearch) ([]model.Employee, error) {
	rs, err := e.adapter.Fortress().GetEmployees(fortress.EmployeeSearch{
		DiscordID: in.DiscordID,
		Email:     in.Email,
		Key:       in.Key,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (e *Profile) GetDiscordRoles(guildID string, userID string) (rs []string, err error) {
	// Get member roles
	gm, err := e.ses.GuildMember(guildID, userID)
	if err != nil {
		return nil, err
	}

	// Get guild to access role positions
	guild, err := e.ses.Guild(guildID)
	if err != nil {
		return nil, err
	}

	// Create a map of role ID to position
	rolePositions := make(map[string]int)
	for _, guildRole := range guild.Roles {
		rolePositions[guildRole.ID] = guildRole.Position
	}

	// Convert member roles to a slice we can sort
	roles := make([]struct {
		ID       string
		Position int
	}, len(gm.Roles))

	for i, roleID := range gm.Roles {
		roles[i] = struct {
			ID       string
			Position int
		}{
			ID:       roleID,
			Position: rolePositions[roleID],
		}
	}

	// Sort roles by position (descending)
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Position > roles[j].Position
	})

	// Extract sorted role IDs
	rs = make([]string, len(roles))
	for i, role := range roles {
		rs[i] = role.ID
	}

	return rs, nil
}

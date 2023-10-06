package profile

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Profile struct {
	ses *discordgo.Session
}

func New(ses *discordgo.Session) Viewer {
	return &Profile{
		ses: ses,
	}
}

func (v *Profile) Help(message *model.DiscordMessage) error {
	content := []string{
		"**?profile @user**・get employee profile",
		"*Example:* `?profile @nam`",
	}

	msg := &discordgo.MessageEmbed{
		Title:       "**Welcome to Fortress Discord bot**",
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(v.ses, message, msg)
}

func (v *Profile) List(original *model.DiscordMessage, employees []model.Employee) error {
	content := ""
	msg := &discordgo.MessageEmbed{}
	if len(employees) == 0 {
		msg.Title = ":warning: **Error Message** :warning: "
		msg.Description = fmt.Sprintf("This discord user was not belong to the organization\n\n")
		return base.SendEmbededMessage(v.ses, original, msg)
	}

	if len(employees) == 1 {
		employee := employees[0]
		targetUser, err := v.ses.GuildMember(original.GuildId, employee.DiscordID)
		if err != nil {
			return err
		}

		avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.webp?size=240", targetUser.User.ID, targetUser.User.Avatar)
		if targetUser.Avatar != "" {
			avatar = fmt.Sprintf("https://cdn.discordapp.com/guilds/%v/users/%v/avatars/%v.webp?size=240", original.GuildId, targetUser.User.ID, targetUser.Avatar)
		}

		if len(targetUser.Roles) > 0 {
			roles, err := v.ses.GuildRoles(v.ses.State.Guilds[0].ID)
			if err != nil {
				return err
			}

			dfRoles := SortRoleByLevel(DwarvesRoles(roles))
			roleStr := ""
			previousLevel := -1
			line := 0
			for _, r := range dfRoles {
				for _, ur := range targetUser.Roles {
					if r.ID == ur {
						if r.Level != previousLevel {
							roleStr += fmt.Sprintf("\n`%v.` ", line)
							previousLevel = r.Level
							line++
						}
						roleStr += fmt.Sprintf("<@&%v> ", r.ID)
					}
				}
			}
			content += fmt.Sprintf("**Roles** %v", strings.TrimSuffix(roleStr, " "))
		}

		gender := ":woman:"
		if employee.Gender == "Male" {
			gender = ":man:"
		}
		userInfo := ""
		userInfo += fmt.Sprintf("%v `Full Name. `**%v**\n", gender, employee.FullName)
		userInfo += fmt.Sprintf(":birthday: `DoB.       `%v\n", employee.Birthday.Format("2006 Jan 02"))

		mmaScores := fmt.Sprintf("%v `Mastery.   `%v\n", constant.GetEmoji("BADGE1"), employee.MmaScore.MasteryScore)
		mmaScores += fmt.Sprintf("%v `Meaning.   `%v\n", constant.GetEmoji("BADGE2"), employee.MmaScore.MeaningScore)
		mmaScores += fmt.Sprintf("%v `Autonomy.  `%v\n", constant.GetEmoji("BADGE3"), employee.MmaScore.AutonomyScore)

		userInfo += "\n**MMA Scores**\n" + mmaScores

		userSkill := ""
		positions := ""
		for _, p := range employee.Positions {
			positions += fmt.Sprintf("%v, ", p.Name)
		}

		stacks := ""
		for _, s := range employee.Stacks {
			stacks += fmt.Sprintf("%v, ", s.Name)
		}

		userSkill += fmt.Sprintf("%v `Position.  `%v\n", constant.GetEmoji("BADGE5"), strings.TrimRight(positions, ", "))
		userSkill += fmt.Sprintf(":jigsaw: `Stacks.    `%v\n", strings.TrimRight(stacks, ", "))

		userInfo += "\n**Skills**\n" + userSkill

		projects := ""
		for _, p := range employee.Projects {
			projects += fmt.Sprintf("%v, ", p.Name)
		}
		userInfo += "\n**Projects:**  " + strings.TrimRight(projects, ", ")

		messageEmbed := []*discordgo.MessageEmbedField{
			{
				Name:   "User Info",
				Value:  userInfo,
				Inline: false,
			},
		}

		msg.Author = &discordgo.MessageEmbedAuthor{
			URL:     "https://fortress.d.foundation/" + employee.Username,
			Name:    targetUser.User.Username,
			IconURL: avatar,
		}
		msg.Fields = messageEmbed
		msg.Description = content
	} else {
		listUsers := ""
		for idx, e := range employees {
			listUsers += fmt.Sprintf("%v. %v - `%v` -`%v`\n", idx+1, e.FullName, e.TeamEmail, e.DiscordID)
		}

		messageEmbed := []*discordgo.MessageEmbedField{
			{
				Name:   "Search Result",
				Value:  listUsers,
				Inline: false,
			},
		}

		msg.Title = "🔎 Team Profile Search"
		msg.Fields = messageEmbed
		msg.Description = "To get user profile detail, please search with: \n`?profile @user`\n`?profile email`\n`?profile discordID`"
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}

func DwarvesRoles(r []*discordgo.Role) []model.DwarvesDiscordRole {
	roleMap := constant.DwarvesRole

	dwarvesRoles := make([]model.DwarvesDiscordRole, 0)
	for _, dRole := range r {
		v, ok := roleMap[dRole.Name]
		if ok {
			dwarvesRoles = append(dwarvesRoles, model.DwarvesDiscordRole{
				ID:    dRole.ID,
				Name:  dRole.Name,
				Level: v,
			})
		}
	}

	return dwarvesRoles
}

func SortRoleByLevel(roles []model.DwarvesDiscordRole) []model.DwarvesDiscordRole {
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Level < roles[j].Level
	})
	return roles
}

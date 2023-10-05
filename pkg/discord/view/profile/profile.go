package profile

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/constant"
	"strings"

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

func (v *Profile) Get(original *model.DiscordMessage, employee *model.Employee) error {
	content := ""

	if employee == nil {
		content += fmt.Sprintf("This discord user was not belong to the organization\n\n")
	}

	targetUser, err := v.ses.GuildMember(original.GuildId, employee.DiscordID)
	if err != nil {
		return err
	}

	avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.webp?size=240", targetUser.User.ID, targetUser.User.Avatar)
	if targetUser.Avatar != "" {
		avatar = fmt.Sprintf("https://cdn.discordapp.com/guilds/%v/users/%v/avatars/%v.webp?size=240", original.GuildId, targetUser.User.ID, targetUser.Avatar)
	}

	if len(targetUser.Roles) > 0 {
		content += fmt.Sprintf(":four_leaf_clover: `Roles.`  <@&%v>", targetUser.Roles[0])
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

	msg := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			URL:     "https://fortress.d.foundation/" + employee.Username,
			Name:    targetUser.User.Username,
			IconURL: avatar,
		},
		Fields:      messageEmbed,
		Description: content,
	}

	return base.SendEmbededMessage(v.ses, original, msg)
}

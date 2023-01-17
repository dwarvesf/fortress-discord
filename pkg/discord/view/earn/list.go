package earn

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Earn) List(channelId string, earns []*model.Earn) error {
	var content string
	for i := range earns {
		earn := earns[i]
		content += fmt.Sprintf("[[%d](%s)] %s\n", earn.Reward, buildEarnUrl(earn.Id), earn.Name)
	}

	msg := &discordgo.MessageEmbed{
		Title:       ":ice_cube: **Dwarves Community Earn** :ice_cube: ",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, channelId, msg)
}

func buildEarnUrl(id string) string {
	return fmt.Sprintf("https://www.notion.so/dwarves/9a5ca08b3312492b9a56cea06431842a?v=f733ece66d81446db452fc4101bdc69d&p=%s", id)
}

package earn

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Earn) List(channelId string, earns []*model.Earn) error {
	earns = []*model.Earn{{Id: "1", Name: "Engagement Log", Reward: 1000}, {Id: "2", Name: "Centralize Dwarves Calendar", Reward: 100}}

	var content string
	for i := range earns {
		content += fmt.Sprintf("[[%d](`https://earn.d.foundation`)] %s\n", earns[i].Reward, earns[i].Name)
	}

	msg := &discordgo.MessageEmbed{
		Title:       ":ice_cube: **Dwarves Community Earn** :ice_cube: ",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, channelId, msg)
}

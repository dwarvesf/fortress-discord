package staff

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Staff) ListDemands(original *model.DiscordMessage, positions []*model.StaffingDemand) error {
	var content string

	// group total
	pos := map[string]int{}
	for i := range positions {
		s := strings.Split(positions[i].Request, ",")
		for ii := range s {
			ss := strings.Split(s[ii], "x")
			if len(ss) >= 2 {
				amountStr := strings.TrimSpace(ss[0])
				amount, err := strconv.Atoi(amountStr)
				if err != nil {
					return err
				}
				position := strings.TrimSpace(ss[1])
				if _, ok := pos[position]; !ok {
					pos[position] = amount
				} else {
					pos[position] += amount
				}
			}
		}
	}

	for i := range positions {
		content += fmt.Sprintf("%d ・ %s - %s\n", i+1, positions[i].Name, positions[i].Request)
	}

	content += "\nTotal:\n"

	for k, v := range pos {
		content += fmt.Sprintf("%v %v \n", v, k)
	}

	msg := &discordgo.MessageEmbed{
		Title:       "<:pepebusiness:885513213687504936> Staffing Demands <:pepebusiness:885513213687504936>\n",
		Description: content,
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

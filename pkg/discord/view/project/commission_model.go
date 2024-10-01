package project

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/shopspring/decimal"
)

func (e *Project) CommissionModels(original *model.DiscordMessage, commissionModel []model.ProjectCommissionModel) error {
	var content string
	for i := range commissionModel {
		v := commissionModel[i]
		content += fmt.Sprintf("`Type.        ` **%s**\n", v.CommissionType)
		content += fmt.Sprintf("`Rate.        ` **%v%%**\n", v.CommissionRate)
		content += fmt.Sprintf("`Beneficiary. ` **%s**\n", v.Beneficiary.FullName)
		content += fmt.Sprintf("`Description. ` **%s**\n", v.Description)

		if v.Sub != nil {
			sub := v.Sub
			content += fmt.Sprintf("`Sale Referral.`\n")
			content += fmt.Sprintf("`      Type.        ` **%s**\n", sub.CommissionType)
			content += fmt.Sprintf("`      Rate.        ` **%v%%**\n", sub.CommissionRate.Mul(v.CommissionRate).Div(decimal.NewFromInt(100)))
			content += fmt.Sprintf("`      Beneficiary. ` **%s**\n", sub.Beneficiary.FullName)
			content += fmt.Sprintf("`      Description. ` **%s**\n", sub.Description)
			content += "\n"
		} else {
			content += "\n"
		}

	}
	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:pepe_ping:1028964391690965012> Commission Models <:pepe_ping:1028964391690965012>"),
		Description: content,
	}

	base.SendEmbededMessage(e.ses, original, msg)
	return nil
}

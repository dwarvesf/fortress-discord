package project

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/leekchan/accounting"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

func (e *Project) PnL(original *model.DiscordMessage, pnls []model.ProjectPnL) error {
	ac := accounting.Accounting{Symbol: "$", Precision: 0}

	// Send third message with project summary as embed
	title := "📊 Project Finance Summary"

	// Construct the message content in a compact format with headers
	var builder strings.Builder
	builder.WriteString("```\n")
	builder.WriteString("Prj  |  Cost  |   Rev  |  \n")
	builder.WriteString("-----|--------|--------|--\n")

	// Populate each project row in a compact format
	for _, pnl := range pnls {
		if len(pnl.Code) > 4 {
			pnl.Code = pnl.Code[:4]
		}
		cost := ac.FormatMoney(pnl.EstimatedCost)
		revenue := ac.FormatMoney(pnl.EstimatedRevenue)
		ratio := pnl.RevenueToCostRatio
		ratioIndicator := getRatioIndicator(ratio)

		// Add formatted row with precise alignment
		builder.WriteString(fmt.Sprintf(
			"%-4s | %7s| %7s|%-1s\n",
			pnl.Code,
			cost,
			revenue,
			ratioIndicator,
		))
	}

	builder.WriteString("```")

	msg := &discordgo.MessageEmbed{
		Title:       title,
		Description: builder.String(),
		Color:       3447003,
	}

	base.SendEmbededMessage(e.ses, original, msg)
	return nil
}

// Helper to get up/down indicator based on ratio
func getRatioIndicator(ratio float64) string {
	if ratio > 1.3 {
		return "🟢"
	}
	return "🔴"
}

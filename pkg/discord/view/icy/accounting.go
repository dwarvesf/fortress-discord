package icy

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/cryptoutils"
)

func (e *Icy) Accounting(original *model.DiscordMessage, icyAccounting *model.IcyAccounting, report *model.SalaryAdvanceReport, total30DaysReward *model.ICYTotalEarned) error {
	if original == nil {
		return nil
	}

	content := make([]string, 0)

	var formattedContractFund string = "Not Found"
	if icyAccounting != nil {
		if fContractFund := cryptoutils.StringBigIntToStringCurrency(icyAccounting.ContractFundInUSDT, icyAccounting.USDT.Decimals); fContractFund != "" {
			formattedContractFund = fContractFund
		}
	}
	content = append(content, fmt.Sprintf("%s `Contract Fund.   `%s **%s USDT**\n", discordEmojiPepeMoney, discordEmojiDollar, formattedContractFund))

	var formattedTotal30DaysICY, formattedTotal30DaysUSD string = "Not Found", "Not Found"
	if total30DaysReward != nil {
		formattedTotal30DaysICY = cryptoutils.StringBigIntToStringCurrency(total30DaysReward.TotalEarnsICY, 0)
		formattedTotal30DaysUSD = cryptoutils.StringBigIntToStringCurrency(total30DaysReward.TotalEarnsUSD, 0)
	}
	content = append(content, fmt.Sprintf("%s **Total 30 Days Reward** - %s `%s ICY` (~ %s `%s USDT`)\n", discordEmojiPepeMoney, discordEmojiIcyToken, formattedTotal30DaysICY, discordEmojiDollar, formattedTotal30DaysUSD))

	var formattedTotalUnpaidICY string = "Not Found"
	if report != nil {
		formattedTotalUnpaidICY = cryptoutils.StringBigIntToStringCurrency(big.NewInt(report.TotalICY).String(), 0)
	}
	content = append(content, fmt.Sprintf("%s **Unpaid Salary Advances** - %s `%s ICY`\n", discordEmojiPepeBusiness, discordEmojiIcyToken, formattedTotalUnpaidICY))

	if len(report.SalaryAdvances) == 0 {
		content = append(content, "\nNo unpaid salary advances\n")
	} else {
		for i, advance := range report.SalaryAdvances {
			formattedICYAmount := cryptoutils.StringBigIntToStringCurrency(big.NewInt(advance.AmountICY).String(), 0)
			content = append(content, fmt.Sprintf("%d. <@%s> - %s `%s ICY`", i, advance.DiscordID, discordEmojiIcyToken, formattedICYAmount))
		}
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s ICY Accounting", discordEmojiPepeNote),
		Description: strings.Join(content, "\n"),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

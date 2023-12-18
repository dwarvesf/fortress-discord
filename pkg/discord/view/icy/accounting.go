package icy

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/dwarvesf/fortress-discord/pkg/discord/view/base"
	"github.com/dwarvesf/fortress-discord/pkg/model"
	"github.com/dwarvesf/fortress-discord/pkg/utils/cryptoutils"
)

const (
	discordEmojiIceCube    = ":ice_cube:"
	discordEmojiContract   = "<:proposal:1087564986504708167>"
	discordEmojiCash       = "<:cash:1058304283642167319>"
	discordEmojiAccounting = "<:nft:1090477914388172840>"
	discordEmojiConversion = "<:conversion:1100681077808443423>"
	discordEmojiShrugging  = "<a:shrugging:1095990328898637824>"
	discordEmojiUSDT       = "<:usdt:1113120155568971786>"
)

func (e *Icy) Accounting(original *model.DiscordMessage, icyAccounting *model.IcyAccounting, report *model.SalaryAdvanceReport) error {
	if original == nil || icyAccounting == nil || icyAccounting.ICY == nil || icyAccounting.USDT == nil || report == nil {
		return nil
	}

	var formattedContractFund, formattedCirculatingICY, formattedOffsetUSDT string = "NaN", "NaN", "NaN"
	if fContractFund := cryptoutils.StringBigIntToStringCurrency(icyAccounting.ContractFundInUSDT, icyAccounting.USDT.Decimals); fContractFund != "" {
		formattedContractFund = fContractFund
	}

	if fCirculatingICY := cryptoutils.StringBigIntToStringCurrency(icyAccounting.CirculatingICY, icyAccounting.ICY.Decimals); fCirculatingICY != "" {
		formattedCirculatingICY = fCirculatingICY
	}

	if fOffsetUSDT := cryptoutils.StringBigIntToStringCurrency(icyAccounting.OffsetUSDT, icyAccounting.USDT.Decimals); fOffsetUSDT != "" {
		formattedOffsetUSDT = fOffsetUSDT
	}

	contentLines := []string{
		fmt.Sprintf("%s `Contract Fund.   `%s **%s USDT**", discordEmojiContract, discordEmojiUSDT, formattedContractFund),
		fmt.Sprintf("%s `Circulating ICY. `%s **%s ICY**", discordEmojiConversion, ":ice_cube:", formattedCirculatingICY),
		fmt.Sprintf("%s `Offset USDT.     `%s **%s USDT**", discordEmojiCash, discordEmojiUSDT, formattedOffsetUSDT),
		"",
		fmt.Sprintf("%s **Unpaid Salary Advances** - %s `%d ICY`", discordEmojiShrugging, discordEmojiIceCube, report.TotalICY),
	}

	for i, advance := range report.SalaryAdvances {
		contentLines = append(contentLines, fmt.Sprintf("%d. <@%s> - %s `%d ICY`", i, advance.DiscordID, discordEmojiIceCube, advance.AmountICY))
	}

	if icyAccounting.ICY.Address != "" {
		icyAddress := strings.TrimSpace(icyAccounting.ICY.Address)
		contentLines = append(contentLines, fmt.Sprintf("\n Fund the contract by transfer USDT to this address `%s` on Polygon", icyAddress))
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s ICY Accounting", discordEmojiAccounting),
		Description: strings.Join(contentLines, "\n"),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

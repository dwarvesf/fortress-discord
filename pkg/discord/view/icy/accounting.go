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
	discordEmojiAccountingIceCube        = ":ice_cube:"
	discordEmojiAccountingDollar         = ":dollar:"
	discordEmojiAccountingTitle          = "<:pepenote:885515949673951282>"
	discordEmojiAccountingContractFund   = "<:pepeMoney:1086173791329210388>"
	discordEmojiAccountingCirculatingICY = "<:anxinicy:1014483263705862174>"
	discordEmojiAccountingOffsetUSDT     = "<:anxin:973799916147179550>"
	discordEmojiAccountingUnpaid         = "<:pepebusiness:885513213687504936>"
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
		fmt.Sprintf("%s `Contract Fund.   `%s **%s USDT**", discordEmojiAccountingContractFund, discordEmojiAccountingDollar, formattedContractFund),
		fmt.Sprintf("%s `Circulating ICY. `%s **%s ICY**", discordEmojiAccountingCirculatingICY, discordEmojiAccountingIceCube, formattedCirculatingICY),
		fmt.Sprintf("%s `Offset USDT.     `%s **%s USDT**", discordEmojiAccountingOffsetUSDT, discordEmojiAccountingDollar, formattedOffsetUSDT),
		"",
		fmt.Sprintf("%s **Unpaid Salary Advances** - %s `%d ICY`", discordEmojiAccountingUnpaid, discordEmojiAccountingIceCube, report.TotalICY),
	}

	for i, advance := range report.SalaryAdvances {
		contentLines = append(contentLines, fmt.Sprintf("%d. <@%s> - %s `%d ICY`", i, advance.DiscordID, discordEmojiAccountingIceCube, advance.AmountICY))
	}

	if icyAccounting.ICY.Address != "" {
		icyAddress := strings.TrimSpace(icyAccounting.ICY.Address)
		contentLines = append(contentLines, fmt.Sprintf("\nFund the contract by transfer USDT to this address `%s` on Polygon", icyAddress))
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s ICY Accounting", discordEmojiAccountingTitle),
		Description: strings.Join(contentLines, "\n"),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

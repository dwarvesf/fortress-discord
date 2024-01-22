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

func (e *Icy) PersonalInfo(original *model.DiscordMessage, accounting *model.IcyAccounting, totalEarned *model.ICYTotalEarned, earnedTxns []*model.ICYEarnedTransaction) error {
	// Do nothing when there is no message to reply to
	if original == nil {
		return nil
	}

	contentLines := make([]string, 0)

	contentLines = append(contentLines, icyAddressContent(accounting)...)
	contentLines = append(contentLines, "")
	contentLines = append(contentLines, conversionRateContent(accounting)...)
	contentLines = append(contentLines, "")
	contentLines = append(contentLines, contractFundContent(accounting)...)
	contentLines = append(contentLines, "")
	contentLines = append(contentLines, yourEarnedICYContent(totalEarned)...)
	contentLines = append(contentLines, "")
	contentLines = append(contentLines, last5TxnsContent(earnedTxns)...)

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s Your ICY Discovery", discordEmojiPepeWhale),
		Description: strings.Join(contentLines, "\n"),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

func icyAddressContent(accounting *model.IcyAccounting) (result []string) {
	result = append(result, fmt.Sprintf("%s **ICY Address**", discordEmojiPepeStonk))

	if accounting != nil && accounting.ICY != nil {
		result = append(result, fmt.Sprintf("`%s|%s`", strings.ToUpper(accounting.ICY.Chain), accounting.ICY.Address))
	} else {
		result = append(result, "Not found")
	}

	return
}

func conversionRateContent(accounting *model.IcyAccounting) (result []string) {
	result = append(result, fmt.Sprintf("%s **Conversion Rate**", discordEmojiPepeAhegao))

	if accounting != nil {
		result = append(result, fmt.Sprintf("%s`1 ICY` ~ %s`%.1f USDT`", discordEmojiIcyToken, discordEmojiDollar, accounting.ConversionRate))
	} else {
		result = append(result, "Not found")
	}

	return
}

func contractFundContent(accounting *model.IcyAccounting) (result []string) {
	result = append(result, fmt.Sprintf("%s **Contract Fund**", discordEmojiPepeMoney))

	if accounting != nil && accounting.USDT != nil {
		if fund := cryptoutils.StringBigIntToStringCurrency(accounting.ContractFundInUSDT, accounting.USDT.Decimals); fund != "" {
			result = append(result, fmt.Sprintf("%s`%s USDT`", discordEmojiDollar, fund))
		}
	}

	if len(result) == 1 {
		result = append(result, "Not found")
	}

	return
}

func yourEarnedICYContent(totalEarned *model.ICYTotalEarned) (result []string) {
	result = append(result, fmt.Sprintf("%s **Your Earned ICY**", discordEmojiPepeCoolNerd))

	var (
		formattedTotalEarnedICY  string = "NaN"
		formattedTotalEarnedUSDT string = "NaN"
	)
	if totalEarned != nil {
		if totalEarnedICY := cryptoutils.StringBigIntToStringCurrency(totalEarned.TotalEarnsICY, 0); totalEarnedICY != "" {
			formattedTotalEarnedICY = totalEarnedICY
		}
		if totalEarnedUSDT := cryptoutils.StringBigIntToStringCurrency(big.NewFloat(totalEarned.TotalEarnsUSD).String(), 0); totalEarnedUSDT != "" {
			formattedTotalEarnedUSDT = totalEarnedUSDT
		}
	}

	result = append(result, fmt.Sprintf("%s`%s ICY` ~ %s`%s USDT`", discordEmojiIcyToken, formattedTotalEarnedICY, discordEmojiDollar, formattedTotalEarnedUSDT))
	return
}

func last5TxnsContent(earnedTxns []*model.ICYEarnedTransaction) (result []string) {
	result = append(result, fmt.Sprintf("%s **Your Last 5 Earns**", discordEmojiAnxinICY))

	if len(earnedTxns) == 0 {
		result = append(result, "No transaction found!")
		return
	}

	last5Txns := [][]string{}
	last5Msgs := []string{}
	for _, txn := range earnedTxns {
		if txn == nil || txn.Token == nil {
			continue
		}

		var amountICY string = "NaN"
		if amount := cryptoutils.StringBigIntToStringCurrency(txn.Amount, int(txn.Token.Decimal)); amount != "" {
			amountICY = amount
		}

		var amountUSDT string = "NaN"
		if amount := cryptoutils.StringBigIntToStringCurrency(big.NewFloat(txn.USDAmount).String(), 0); amount != "" {
			amountUSDT = amount
		}

		msg, _ := txn.Metadata["message"].(string)
		last5Msgs = append(last5Msgs, msg)

		last5Txns = append(last5Txns, []string{
			formatTimeAgo(txn.CreatedAt),
			txn.ExternalID[:4],
			fmt.Sprintf("%s ICY", amountICY),
			fmt.Sprintf("%s USDT", amountUSDT),
		})
	}

	result = append(result, formatTable(last5Txns, last5Msgs))
	return
}

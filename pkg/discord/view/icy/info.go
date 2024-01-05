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
	if original == nil || accounting == nil || accounting.ICY == nil || accounting.USDT == nil || totalEarned == nil || earnedTxns == nil {
		return nil
	}

	var formattedContractFund string = "NaN"
	if fund := cryptoutils.StringBigIntToStringCurrency(accounting.ContractFundInUSDT, accounting.USDT.Decimals); fund != "" {
		formattedContractFund = fund
	}

	var formattedTotalEarnedICY string = "NaN"
	if totalEarnedICY := cryptoutils.StringBigIntToStringCurrency(totalEarned.TotalEarnsICY, 0); totalEarnedICY != "" {
		formattedTotalEarnedICY = totalEarnedICY
	}

	var formattedTotalEarnedUSDT string = "NaN"
	if totalEarnedUSDT := cryptoutils.StringBigIntToStringCurrency(big.NewFloat(totalEarned.TotalEarnsUSD).String(), 0); totalEarnedUSDT != "" {
		formattedTotalEarnedUSDT = totalEarnedUSDT
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

	txnsTable := formatTable(last5Txns, last5Msgs)

	contentLines := []string{
		// ICY Address
		fmt.Sprintf("%s **ICY Address**", discordEmojiPepeStonk),
		fmt.Sprintf("`%s|%s`", strings.ToUpper(accounting.ICY.Chain), accounting.ICY.Address),
		"",
		// ICY conversion rate
		fmt.Sprintf("%s **Conversion Rate**", discordEmojiPepeAhegao),
		fmt.Sprintf("%s`1 ICY` ~ %s`%.1f USDT`", discordEmojiIcyToken, discordEmojiDollar, accounting.ConversionRate),
		"",
		// Contract Fund
		fmt.Sprintf("%s **Contract Fund**", discordEmojiPepeMoney),
		fmt.Sprintf("%s`%s USDT`", discordEmojiDollar, formattedContractFund),
		"",
		// Your Earned ICY
		fmt.Sprintf("%s **Your Earned ICY**", discordEmojiPepeCoolNerd),
		fmt.Sprintf("%s`%s ICY` ~ %s`%s USDT`", discordEmojiIcyToken, formattedTotalEarnedICY, discordEmojiDollar, formattedTotalEarnedUSDT),
		"",
		// Your Last 5 Transactions
		fmt.Sprintf("%s **Your Last 5 Earns**", discordEmojiAnxinICY),
		txnsTable,
	}

	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s Your ICY Discovery", discordEmojiPepeWhale),
		Description: strings.Join(contentLines, "\n"),
	}

	return base.SendEmbededMessage(e.ses, original, msg)
}

package discord

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/bwmarrin/discordgo"
	fm "github.com/consolelabs/mochi-toolkit/formatter"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (d *Discord) SetStatus() error {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	for {
		err := d.updateStatus()
		if err != nil {
			d.L.Error(err, "failed to set status")
		}

		<-ticker.C
	}
}

func (d *Discord) updateStatus() error {
	// Get total volume ICY issue
	data, err := d.Svc.Tono().GetCommunityTransaction()
	if err != nil {
		return fmt.Errorf("failed to get community transaction: %w", err)
	}

	// Get USDC balance
	balance, err := getUSDCBalance()
	if err != nil {
		return fmt.Errorf("failed to get USDC balance: %w", err)
	}

	// Create status messages
	statusMessages := []string{
		fmt.Sprintf("Contract Fund • %s USDC", fm.FormatTokenAmount(balance.String(), 6)),
		fmt.Sprintf("Issued ICY • %2.f ICY", data.TotalRewardVolume),
	}

	// Rotate through status messages
	currentTime := time.Now()
	index := currentTime.Minute() / 30 % len(statusMessages)

	err = d.Session.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name:  statusMessages[index],
				Type:  discordgo.ActivityTypeCustom,
				State: statusMessages[index],
			},
		},
		AFK:    false,
		Status: "",
	})
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}

func getUSDCBalance() (*big.Int, error) {
	client, err := ethclient.Dial("https://mainnet.base.org")
	if err != nil {
		log.Fatalf("Failed to connect to the Base chain: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x982d2c5A654E4f7CC65ACDCa4ECc649fE4F4DAa4")

	usdtAddress := common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913")

	data := common.Hex2Bytes("70a08231000000000000000000000000" + contractAddress.Hex()[2:])

	ctx := context.Background()
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &usdtAddress,
		Data: data,
	}, nil)

	if err != nil {
		return nil, err
	}

	balance := new(big.Int).SetBytes(result)
	return balance, nil
}

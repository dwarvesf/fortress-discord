package model

import "time"

// AdapterIcy is a struct response from adapter, before process to in-app model
type AdapterIcy struct {
	Data    []*Icy
	Message string
}

// Icy is in-app model, after process from adapters
type Icy struct {
	ID     string `json:"id"`
	Period string `json:"period"`
	Team   string `json:"team"`
	Amount string `json:"amount"`
}

type AdapterIcyAccounting struct {
	Data *IcyAccounting `json:"data"`
}

type IcyAccounting struct {
	ICY     *TokenInfo    `json:"icy"`
	USDT    *TokenInfo    `json:"usdt"`
	IcySwap *ContractInfo `json:"icySwap"`

	ConversionRate     float32 `json:"conversionRate"`
	CirculatingICY     string  `json:"circulatingICY"`
	ContractFundInUSDT string  `json:"contractFundInUSDT"`
	OffsetUSDT         string  `json:"offsetUSDT"` // how many usdt left need to be issued
}

type TokenInfo struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Address     string `json:"address"`
	Decimals    int    `json:"decimals"`
	Chain       string `json:"chain"`
	ChainID     string `json:"chainID"`
	TotalSupply string `json:"totalSupply"`
}

type ContractInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Chain   string `json:"chain"`
}

type AdapterICYEarnedTransactions struct {
	Data []*ICYEarnedTransaction `json:"data"`
}

type ICYEarnedTransaction struct {
	ID                 string                 `json:"id"`
	FromProfileID      string                 `json:"fromProfileID"`
	OtherProfileID     string                 `json:"otherProfileID"`
	FromProfileSource  string                 `json:"fromProfileSource"`
	OtherProfileSource string                 `json:"otherProfileSource"`
	SourcePlatform     string                 `json:"sourcePlatform"`
	Amount             string                 `json:"amount"`
	TokenID            string                 `json:"tokenID"`
	ChainID            string                 `json:"chainID"`
	InternalID         int64                  `json:"internalID"`
	ExternalID         string                 `json:"externalID"`
	OnchainTxHash      string                 `json:"onchainTxHash"`
	Type               string                 `json:"type"`
	Action             string                 `json:"action"`
	Status             string                 `json:"status"`
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time              `json:"updatedAt"`
	ExpiredAt          *time.Time             `json:"expiredAt"`
	SettledAt          *time.Time             `json:"settledAt"`
	Token              *MochiToken            `json:"token"`
	OriginalTxID       string                 `json:"originalTxID"`
	OtherProfile       *MochiProfile          `json:"otherProfile"`
	FromProfile        *MochiProfile          `json:"fromProfile"`
	USDAmount          float64                `json:"usdAmount"`
	Metadata           map[string]interface{} `json:"metadata"`
	OtherProfileIds    []string               `json:"otherProfileIds"`
	TotalAmount        string                 `json:"totalAmount"`
	FromTokenId        string                 `json:"fromTokenId"`
	ToTokenId          string                 `json:"toTokenId"`
	FromAmount         string                 `json:"fromAmount"`
	ToAmount           string                 `json:"toAmount"`
}

type AdapterICYTotalEarned struct {
	Data *ICYTotalEarned `json:"data"`
}

type ICYTotalEarned struct {
	TotalEarnsICY string `json:"totalEarnsICY"`
	TotalEarnsUSD string `json:"totalEarnsUSD"`
}

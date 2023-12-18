package model

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

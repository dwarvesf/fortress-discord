package model

type MochiProfile struct {
	ID                 string                    `json:"id"`
	CreatedAt          string                    `json:"createdAt"`
	UpdatedAt          string                    `json:"updatedAt"`
	ProfileName        string                    `json:"profileName"`
	Avatar             string                    `json:"avatar"`
	AssociatedAccounts []MochiAssociatedAccounts `json:"associatedAccounts"`
	Type               string                    `json:"type"`
}

type MochiAssociatedAccounts struct {
	ID                 string      `json:"id"`
	ProfileID          string      `json:"profileID"`
	Platform           string      `json:"platform"`
	PlatformIdentifier string      `json:"platformIdentifier"`
	PlatformMetadata   interface{} `json:"platformMetadata"`
	IsGuildMember      bool        `json:"isGuildMember"`
	CreatedAt          string      `json:"createdAt"`
	UpdatedAt          string      `json:"updatedAt"`
}

type MochiToken struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Decimal     int64   `json:"decimal"`
	ChainID     string  `json:"chainID"`
	Native      bool    `json:"native"`
	Address     string  `json:"address"`
	Icon        string  `json:"icon"`
	CoinGeckoID string  `json:"coinGeckoID"`
	Price       float64 `json:"price"`
}

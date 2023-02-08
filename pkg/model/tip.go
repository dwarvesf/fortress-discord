package model

type Tip struct {
	Recipients   []string `json:"recipients"`
	Sender       string   `json:"sender"`
	GuildId      string   `json:"guild_id"`
	ChannelId    string   `json:"channel_id"`
	Amount       float64  `json:"amount"`
	Token        string   `json:"token"`
	TransferType string   `json:"transfer_type"`
	Platform     string   `json:"platform"`
}

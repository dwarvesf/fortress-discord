package model

type ListGuildCommunityTransaction struct {
	TotalVolume           string  `json:"total_volume"`
	TotalVolumeInUsd      string  `json:"total_volume_in_usd"`
	TopReceiver           string  `json:"top_receiver"`
	TotalTransactions     int64   `json:"total_transactions"`
	TotalRewardTxs        int64   `json:"total_reward_txs"`
	TotalMemberTxs        int64   `json:"total_member_txs"`
	TotalRewardVolume     float64 `json:"total_reward_volume"`
	AvgPerTx              string  `json:"avg_per_tx"`
	TotalMemberReceivedTx int     `json:"total_member_received_tx"`
	TxChange              float64 `json:"tx_change"`
	VolumeChange          float64 `json:"volume_change"`
	NewUser               int     `json:"new_user"`
}

type ListGuildCommunityTransactionResponse struct {
	Data ListGuildCommunityTransaction `json:"data"`
}

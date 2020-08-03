package binance

import "time"


type (
	NodeInfoResponse struct {
		SyncInfo struct {
			LatestBlockHeight int `json:"latest_block_height"`
		} `json:"sync_info"`
	}

	TransactionsInBlockResponse struct {
		BlockHeight int `json:"blockHeight"`
		Tx          []struct {
			TxHash      string      `json:"txHash"`
			BlockHeight int         `json:"blockHeight"`
			TxType      string      `json:"txType"`
			TimeStamp   time.Time   `json:"timeStamp"`
			FromAddr    string      `json:"fromAddr"`
			ToAddr      interface{} `json:"toAddr"`
			Value       string      `json:"value"`
			TxAsset     string      `json:"txAsset"`
			TxFee       string      `json:"txFee"`
			OrderID     string      `json:"orderId,omitempty"`
			Code        int         `json:"code"`
			Data        string      `json:"data"`
			Memo        string      `json:"memo"`
			Source      int         `json:"source"`
			Sequence    int         `json:"sequence"`
		} `json:"tx"`
	}

	TxType string
)

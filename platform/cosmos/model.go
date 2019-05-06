package cosmos

type Tx struct {
	BlockHeight uint64 `json:"height"`
	Timestamp   string `json:"timestamp"`
	Hash        string `json:"txhash"`
	TxData      struct {
		TxType  string `json:"type"`
		TxValue struct {
			TxMessage struct {
				TxContents []TxContents `json:"value"`
			} `json:"msg"`
			TxFee  []TxFee `json:"fee"`
			TxMemo string  `json:"memo"`
		} `json:"value"`
	} `json:"tx"`
}

type TxContents struct {
	FromAddr string   `json:"from_address"`
	ToAddr   string   `json:"to_address"`
	TxAmount []Amount `json:"amount"`
}

type TxFee struct {
	FeeAmount []Amount `json:"amount"`
}

type Amount struct {
	Denom  string `json:"denom"`
	Amount uint64 `json:"amount"`
}

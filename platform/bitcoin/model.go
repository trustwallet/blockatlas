package bitcoin

type TransferTx struct {
	Transactions []string `json:"transactions"`
}

type XpubTransfers struct {
	Transactions []TransferReceipt `json:"transactions"`
}

type Tx struct {
	ID string `json:"id"`
}

type TransferReceipt struct {
	ID            string     `json:"txid"`
	Version       uint64     `json:"version"`
	Vin           []Transfer `json:"vin"`
	Vout          []Transfer `json:"vout"`
	BlockHash     string     `json:"blockHash"`
	BlockHeight   uint64     `json:"blockHeight"`
	Confirmations uint64     `json:"confirmations"`
	BlockTime     uint64     `json:"blockTime"`
	Value         string     `json:"value"`
	ValueIn       string     `json:"valueIn"`
	Fees          string     `json:"fees"`
	Hex           string     `json:"hex"`
}

type Transfer struct {
	TxId      string   `json:"txid"`
	Sequence  uint64   `json:"sequence"`
	Value     string   `json:"value"`
	Addresses []string `json:"addresses"`
	Hex       string   `json:"hex"`
}

type Meta struct {
	BlockID        string `json:"blockID"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
	TxID           string `json:"txID"`
	TxOrigin       string `json:"txOrigin"`
}

type TransferAddress struct {
	Address string `json:"address"`
}

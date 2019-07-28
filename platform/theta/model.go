package theta

// THETA transaction types https://github.com/thetatoken/theta-mainnet-integration-guide/blob/master/docs/api.md#getblock
const (
	SendTransaction = 2
)

// Response from Explorer
type AccountTxList struct {
	Body []Tx `json:"body"`
}

type Tx struct {
	Hash        string `json:"hash"`
	Type        int    `json:"type"`
	Data        Data   `json:"data"`
	BlockHeight string `json:"block_height"`
	Timestamp   string `json:"timestamp"`
}

type Data struct {
	Fee     Fee      `json:"fee"`
	Inputs  []Inputs `json:"inputs"`
	Outputs []Output `json:"outputs"`
}

type Fee struct {
	Thetawei string `json:"thetawei"`
	Tfuelwei string `json:"tfuelwei"`
}

type Inputs struct {
	Address  string `json:"address"`
	Sequence string `json:"sequence"`
}
type Output struct {
	Address string `json:"address"`
	Coins   Fee    `json:"coins"`
}

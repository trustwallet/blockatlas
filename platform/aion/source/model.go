package source

type TxPage struct {
	Content []Tx
} 

type Tx struct {
	BlockHash            string  `json:"blockHash"`
	ToAddr               string  `json:"toAddr"`
	ContractAddr         string  `json:"contractAddr"`
	TransactionHash      string  `json:"transactionHash"`
	TransactionTimestamp int64   `json:"transactionTimestamp"`
	NrgConsumed          int     `json:"nrgConsumed"`
	BlockNumber          uint64  `json:"blockNumber"`
	BlockTimestamp       int64   `json:"blockTimestamp"`
	FromAddr             string  `json:"fromAddr"`
	Value                float64 `json:"value"`
}

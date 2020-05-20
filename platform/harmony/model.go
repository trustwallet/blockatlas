package harmony

type TxResponse struct {
	Result TxResult `json:"result"`
}

type TxResult struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Hash        string `json:"hash"`
	Nonce       string `json:"nonce"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Timestamp   string `json:"timestamp"`
}

type BlockInfo struct {
	Hash         string        `json:"hash"`
	Number       string        `json:"number"`
	Transactions []Transaction `json:"transactions"`
}

type ValidatorInfo struct {
	Address string  `json:"address"`
}

type LifetimeInfo struct {
	Apr     string `json:"apr"`
}

type Validator struct {
	Info      ValidatorInfo    `json:"validator"`
	Active    bool             `json:"currently-in-committee"`
	Lifetime  LifetimeInfo      `json:"lifetime"`
}

type Validators struct {
	Validators []Validator `json:"result"`
}

type Delegation struct {
	DelegatorAddress string   `json:"delegator_address"`
	ValidatorAddress string   `json:"validator_address"`
	Amount           float64  `json:"amount"`
}

type Delegations struct {
	List []Delegation `json:"result"`
}

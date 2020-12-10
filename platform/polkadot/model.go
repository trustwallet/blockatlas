package polkadot

const (
	FeeTransfer string = "100000000"

	ModuleBalances string = "balances"
	ModuleStaking  string = "staking"

	ModuleFunctionTransfer string = "transfer"
)

type Transfer struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Module      string `json:"module"`
	Amount      string `json:"amount"`
	Hash        string `json:"hash"`
	Timestamp   uint64 `json:"block_timestamp"`
	BlockNumber uint64 `json:"block_num"`
	Success     bool   `json:"success"`
}

type Extrinsic struct {
	Timestamp          uint64 `json:"block_timestamp"`
	BlockNumber        uint64 `json:"block_num"`
	CallModuleFunction string `json:"call_module_function"`
	CallModule         string `json:"call_module"`
	Params             string `json:"params"`
	AccountId          string `json:"account_id"`
	Nonce              uint64 `json:"nonce"`
	Hash               string `json:"extrinsic_hash"`
	Success            bool   `json:"success"`
	Fee                string `json:"fee"`
}

type CallData struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TransfersRequest struct {
	Address string `json:"address"`
	Row     int    `json:"row"`
}

type BlockRequest struct {
	BlockNumber int64 `json:"block_num"`
}

type SubscanResponseData struct {
	BlockNumber string      `json:"blockNum,omitempty"`
	Transfers   []Transfer  `json:"transfers,omitempty"`
	Extrinsics  []Extrinsic `json:"extrinsics,omitempty"`
}

type SubscanResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    SubscanResponseData `json:"data"`
}

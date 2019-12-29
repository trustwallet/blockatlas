package polkadot

type Metadata struct {
	BlockNum    string `json:"blockNum"`
	BlockTime   string `json:"blockTime"`
	ImplName    string `json:"implName"`
	NetworkNode string `json:"networkNode"`
	SpecVersion string `json:"specVersion"`
}

type Transfer struct {
	From           string `json:"from"`
	To             string `json:"to"`
	Module         string `json:"module"`
	Amount         string `json:"amount"`
	Hash           string `json:"hash"`
	BlockTimestamp uint64 `json:"block_timestamp"`
	BlockNum       uint64 `json:"block_num"`
	Success        bool   `json:"success"`
}

type Extrinsic struct {
	BlockTimestamp     uint64 `json:"block_timestamp"`
	BlockNum           uint64 `json:"block_num"`
	ValueRaw           string `json:"value_raw"`
	CallModuleFunction string `json:"call_module_function"`
	CallModule         string `json:"call_module"`
	Params             string `json:"params"`
	AccountID          string `json:"account_id"`
	Nonce              uint64 `json:"nonce"`
	Era                string `json:"era"`
	ExtrinsicHash      string `json:"extrinsic_hash"`
	Success            bool   `json:"success"`
}

type TransferRequest struct {
	Address string `json:"address"`
	Row     int    `json:"row"`
}

type SubscanResponseData struct {
	Count      int         `json:"count"`
	Transfers  []Transfer  `json:"transfers,omitempty"`
	Extrinsics []Extrinsic `json:"extrinsics,omitempty"`
}

type SubscanResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    interface{}
}

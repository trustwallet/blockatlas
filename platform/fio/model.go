package fio

// ActionData (from get_actions)
type ActionData struct {
	From           string `json:"from"`
	To             string `json:"to"`
	PayeePublicKey string `json:"payee_public_key"`
	Amount 		   int64  `json:"amount"`
	Quantity       string `json:"quantity"`
	Fee            int64  `json:"max_fee"`
	Actor          string `json:"actor"`
	TpID           string `json:"tpid"`
	Memo    string     `json:"memo"`
}
// ActionAct (from get_actions)
type ActionAct struct {
	Account string     `json:"account"`
	Name 	string     `json:"name"`
	Data    ActionData `json:"data"`
}
// ActionTrace
type ActionTrace struct {
	Receiver string `json:"receiver"`
	Act 	 ActionAct
	TrxID 	 string `json:"trx_id"`
	BlockNum uint64 `json:"block_num"`
}
// Action (from get_actions)
type Action struct {
	BlockNum 	uint64      `json:"block_num"`
	ActionTrace ActionTrace `json:"action_trace"`
}

// GetActionsRequest request struct for get_actions
type GetActionsRequest struct {
	AccountName string `json:"account_name"`
}

// GetActionsResponse request struct for get_actions
type GetActionsResponse struct {
	Actions []Action `json:"actions"`
	//Error 	string	  `json:"error"`
}

// GetPubAddressRequest request struct for get_pub_address
type GetPubAddressRequest struct {
	FioAddress string `json:"fio_address"`
	TokenCode  string `json:"token_code"`
	ChainCode  string `json:"chain_code"`
}

// GetPubAddressResponse response struct for get_pub_address
type GetPubAddressResponse struct {
	PublicAddress string `json:"public_address"`
	BlockNum      int    `json:"block_num"`
	Message       string `json:"message"`
}

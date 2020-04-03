package fio

// ActionData (from get_actions)
type ActionData struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Memo     string `json:"memo"`
}

// ActionAct (from get_actions)
type ActionAct struct {
	Account string      `json:"account"`
	Name    string      `json:"name"`
	Data    interface{} `json:"data"` // Structure of data is action-specific
}

// ActionTrace
type ActionTrace struct {
	Act   ActionAct `json:"act"`
	TrxID string    `json:"trx_id"`
}

// Action (from get_actions)
type Action struct {
	BlockNum    uint64      `json:"block_num"`
	BlockTime   string      `json:"block_time"`
	ActionTrace ActionTrace `json:"action_trace"`
}

// GetActionsRequest request struct for get_actions
// see https://github.com/EOSIO/eos/blob/master/plugins/history_plugin/include/eosio/history_plugin/history_plugin.hpp
type GetActionsRequest struct {
	AccountName string `json:"account_name"`
	Pos         int32  `json:"pos"`
	Offset      int32  `json:"offset"`
	Sort        string `json:"sort"` // desc
}

// GetActionsResponse request struct for get_actions
type GetActionsResponse struct {
	Actions []Action `json:"actions"`
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
	Message       string `json:"message"`
}

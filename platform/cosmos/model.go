package cosmos

import (
	"encoding/json"
)

const (
	CosmosMsgSend                        = "cosmos-sdk/MsgSend"
	CosmosMsgMultiSend                   = "cosmos-sdk/MsgMultiSend"
	CosmosMsgCreateValidator             = "cosmos-sdk/MsgCreateValidator"
	CosmosMsgDelegate                    = "cosmos-sdk/MsgDelegate"
	CosmosMsgUndelegate                  = "cosmos-sdk/MsgUndelegate"
	CosmosMsgBeginRedelegate             = "cosmos-sdk/MsgBeginRedelegate"
	CosmosMsgWithdrawDelegationReward    = "cosmos-sdk/MsgWithdrawDelegationReward"
	CosmosMsgWithdrawValidatorCommission = "cosmos-sdk/MsgWithdrawValidatorCommission"
	CosmosMsgSubmitProposal              = "cosmos-sdk/MsgSubmitProposal"
	CosmosMsgDeposit                     = "cosmos-sdk/MsgDeposit"
	CosmosMsgVote                        = "cosmos-sdk/MsgVote"
	CosmosTextProposal                   = "cosmos-sdk/TextProposal"
	CosmosMsgUnjail                      = "cosmos-sdk/MsgUnjail"
)

// Tx - Base transaction object. Always returned as part of an array
type Tx struct {
	Block string `json:"height"`
	Date  string `json:"timestamp"`
	ID    string `json:"txhash"`
	Data  Data   `json:"tx"`
}

// Data - "tx" sub object
type Data struct {
	Contents Contents `json:"value"`
}

// Contents - amount, fee, and memo
type Contents struct {
	Message []Message `json:"msg"`
	Fee     Fee       `json:"fee"`
	Memo    string    `json:"memo"`
}

// Message - an array that holds multiple 'particulars' entries. Possibly used for multiple transfers in one transaction?
type Message struct {
	Type  string
	Value interface{}
}

// MessageValueTransfer - from, to, and amount
type MessageValueTransfer struct {
	FromAddr string   `json:"from_address"`
	ToAddr   string   `json:"to_address"`
	Amount   []Amount `json:"amount,omitempty"`
}

// MessageValueTransfer - from, to, and amount
type MessageValueDelegate struct {
	DelegatorAddr string `json:"delegator_address"`
	ValidatorAddr string `json:"validator_address"`
	Amount        Amount `json:"amount"`
}

// Fee - also references the "amount" struct
type Fee struct {
	FeeAmount []Amount `json:"amount"`
}

// Amount - the asset & quantity. Always seems to be enclosed in an array/list for some reason.
// Perhaps used for multiple tokens transferred in a single sender/reciever transfer?
type Amount struct {
	Denom    string `json:"denom"`
	Quantity string `json:"amount"`
}

// # Staking

type CosmosValidator struct {
	Status           int                        `json:"status"`
	Description      CosmosValidatorDescription `json:"description"`
	Operator_Address string                     `json:"operator_address"`
	Consensus_Pubkey string                     `json:"consensus_pubkey"`
}

type CosmosValidatorDescription struct {
	Moniker     string `json:"moniker"`
	Website     string `json:"website"`
	Description string `json:"details"`
}

type Block struct {
	Meta BlockMeta `json:"block_meta"`
}

type BlockMeta struct {
	Header BlockHeader `json:"header"`
}

type BlockHeader struct {
	Height string `json:"height"`
}

func (m *Message) UnmarshalJSON(buf []byte) error {
	var messageInternal struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}

	err := json.Unmarshal(buf, &messageInternal)
	if err != nil {
		return err
	}

	m.Type = messageInternal.Type

	switch messageInternal.Type {
	case CosmosMsgUndelegate, CosmosMsgDelegate:
		var msgDelegate MessageValueDelegate
		err = json.Unmarshal(messageInternal.Value, &msgDelegate)
		m.Value = msgDelegate
	case CosmosMsgSend:
		var msgTransfer MessageValueTransfer
		err = json.Unmarshal(messageInternal.Value, &msgTransfer)
		m.Value = msgTransfer
	}
	return err
}

package cosmos

import (
	"encoding/json"
	"strings"
)

// Types of messages
const (
	MsgSend                        = "cosmos-sdk/MsgSend"
	MsgMultiSend                   = "cosmos-sdk/MsgMultiSend"
	MsgCreateValidator             = "cosmos-sdk/MsgCreateValidator"
	MsgDelegate                    = "cosmos-sdk/MsgDelegate"
	MsgUndelegate                  = "cosmos-sdk/MsgUndelegate"
	MsgBeginRedelegate             = "cosmos-sdk/MsgBeginRedelegate"
	MsgWithdrawDelegationReward    = "cosmos-sdk/MsgWithdrawDelegationReward"
	MsgWithdrawValidatorCommission = "cosmos-sdk/MsgWithdrawValidatorCommission"
	MsgSubmitProposal              = "cosmos-sdk/MsgSubmitProposal"
	MsgDeposit                     = "cosmos-sdk/MsgDeposit"
	MsgVote                        = "cosmos-sdk/MsgVote"
	TextProposal                   = "cosmos-sdk/TextProposal"
	MsgUnjail                      = "cosmos-sdk/MsgUnjail"
)

// Tx - Base transaction object. Always returned as part of an array
type Tx struct {
	Block string `json:"height"`
	Date  string `json:"timestamp"`
	ID    string `json:"txhash"`
	Data  Data   `json:"tx"`
}

type TxPage []Tx

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

// MessageValueDelegate - from, to, and amount
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

type CosmosCommission struct {
	Rate string `json:"rate"`
}

type Validator struct {
	Status     int              `json:"status"`
	Address    string           `json:"operator_address"`
	Commission CosmosCommission `json:"commission"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Shares           string `json:"shares,omitempty"`
}

func (d *Delegation) Value() string {
	shares := strings.Split(d.Shares, ".")
	if len(shares) > 0 {
		return shares[0]
	}
	return d.Shares
}

type UnbondingDelegation struct {
	Delegation
	Entries []UnbondingDelegationEntry `json:"entries"`
}

type UnbondingDelegationEntry struct {
	DelegatorAddress string `json:"creation_height"`
	CompletionTime   string `json:"completion_time"`
	Balance          string `json:"balance"`
}

type StakingPool struct {
	NotBondedTokens string `json:"not_bonded_tokens"`
	BondedTokens    string `json:"bonded_tokens"`
}

// Block - top object of get las block request
type Block struct {
	Meta BlockMeta `json:"block_meta"`
}

//BlockMeta - "Block" sub object
type BlockMeta struct {
	Header BlockHeader `json:"header"`
}

//BlockHeader - "BlockMeta" sub object, height
type BlockHeader struct {
	Height string `json:"height"`
}

//UnmarshalJSON reads different message types
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
	case MsgUndelegate, MsgDelegate:
		var msgDelegate MessageValueDelegate
		err = json.Unmarshal(messageInternal.Value, &msgDelegate)
		m.Value = msgDelegate
	case MsgSend:
		var msgTransfer MessageValueTransfer
		err = json.Unmarshal(messageInternal.Value, &msgTransfer)
		m.Value = msgTransfer
	}
	return err
}

type Account struct {
	Value AccountValue `json:"value"`
}

type AccountValue struct {
	Coins []Balance `json:"coins"`
}

type Balance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

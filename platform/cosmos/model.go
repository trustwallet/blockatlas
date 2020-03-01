package cosmos

import (
	"encoding/json"
	"strconv"
	"strings"
)

type TxType string
type EventType string
type AttributeKey string
type DenomType string

// Types of messages
const (
	MsgSend                        TxType = "cosmos-sdk/MsgSend"
	MsgMultiSend                   TxType = "cosmos-sdk/MsgMultiSend"
	MsgCreateValidator             TxType = "cosmos-sdk/MsgCreateValidator"
	MsgDelegate                    TxType = "cosmos-sdk/MsgDelegate"
	MsgUndelegate                  TxType = "cosmos-sdk/MsgUndelegate"
	MsgBeginRedelegate             TxType = "cosmos-sdk/MsgBeginRedelegate"
	MsgWithdrawDelegationReward    TxType = "cosmos-sdk/MsgWithdrawDelegationReward"
	MsgWithdrawValidatorCommission TxType = "cosmos-sdk/MsgWithdrawValidatorCommission"
	MsgSubmitProposal              TxType = "cosmos-sdk/MsgSubmitProposal"
	MsgDeposit                     TxType = "cosmos-sdk/MsgDeposit"
	MsgVote                        TxType = "cosmos-sdk/MsgVote"
	TextProposal                   TxType = "cosmos-sdk/TextProposal"
	MsgUnjail                      TxType = "cosmos-sdk/MsgUnjail"

	EventTransfer        EventType = "transfer"
	EventWithdrawRewards EventType = "withdraw_rewards"

	AttributeAmount    AttributeKey = "amount"
	AttributeValidator AttributeKey = "validator"

	DenomAtom DenomType = "uatom"
	DenomKava DenomType = "ukava"
)

// Tx - Base transaction object. Always returned as part of an array
type Tx struct {
	Block  string `json:"height"`
	Code   int    `json:"code"`
	Date   string `json:"timestamp"`
	ID     string `json:"txhash"`
	Data   Data   `json:"tx"`
	Events Events `json:"events"`
}

type TxPage struct {
	PageTotal string `json:"page_total"`
	Txs       []Tx   `json:"txs"`
}

// Events
type Event struct {
	Type       EventType
	Attributes Attributes `json:"Attributes"`
}

type Events []*Event

func (e Events) GetWithdrawRewardValue() string {
	result := int64(0)
	for _, att := range e {
		if att.Type == EventWithdrawRewards {
			result += att.Attributes.GetWithdrawRewardValue()
		}
	}
	return strconv.FormatInt(result, 10)
}

type Attribute struct {
	Key   AttributeKey `json:"key"`
	Value string       `json:"value"`
}

type Attributes []Attribute

func (a Attributes) GetWithdrawRewardValue() int64 {
	result := int64(0)
	for _, att := range a {
		if att.Key == AttributeAmount {
			idx := strings.IndexByte(att.Value, 'u')
			if idx < 0 {
				continue
			}
			value := att.Value[:idx]
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				continue
			}
			result += v
		}
	}
	return result
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
	Type  TxType
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
	Amount        Amount `json:"amount,omitempty"`
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
	Commision CosmosCommissionRates `json:"commission_rates"`
}

type CosmosCommissionRates struct {
	Rate string `json:"rate"`
}

type Validators struct {
	Result []Validator `json:"result"`
}

type Validator struct {
	Status     int              `json:"status"`
	Address    string           `json:"operator_address"`
	Commission CosmosCommission `json:"commission"`
}

type Inflation struct {
	Result string `json:"result"`
}

type Delegations struct {
	List []Delegation `json:"result"`
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

type UnbondingDelegations struct {
	List []UnbondingDelegation `json:"result"`
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
	Pool Pool `json:"result"`
}

type Pool struct {
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
		Type  TxType          `json:"type"`
		Value json.RawMessage `json:"value"`
	}

	err := json.Unmarshal(buf, &messageInternal)
	if err != nil {
		return err
	}

	m.Type = messageInternal.Type

	switch messageInternal.Type {
	case MsgUndelegate, MsgDelegate, MsgWithdrawDelegationReward:
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

type AuthAccount struct {
	Account Account `json:"result"`
}

type Account struct {
	Value AccountValue `json:"value"`
}

type AccountValue struct {
	Coins []Balance `json:"coins"`
}

type Balance struct {
	Denom  DenomType `json:"denom"`
	Amount string    `json:"amount"`
}

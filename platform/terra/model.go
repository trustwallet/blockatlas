package terra

import (
	"encoding/json"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// TxType nolint
type TxType string

// EventType nolint
type EventType string

// AttributeKey nolint
type AttributeKey string

// DenomType nolint
type DenomType string

// Types of messages
const (
	MsgSend TxType = "bank/MsgSend"

	EventTransfer EventType = "transfer"

	DenomLuna DenomType = "uluna"
)

// Mapping info for internal and external denoms
var (
	DenomMap = map[string]string{
		"uluna": "LUNA",
		"ukrw":  "KRT",
		"usdr":  "SDT",
		"uusd":  "UST",
		"umnt":  "MNT",
	}
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

// TxPage noling
type TxPage struct {
	Txs []Tx `json:"txs"`
}

// Event nolint
type Event struct {
	Type       EventType
	Attributes Attributes `json:"Attributes"`
}

// Events nolint
type Events []*Event

// Attribute nolint
type Attribute struct {
	Key   AttributeKey `json:"key"`
	Value string       `json:"value"`
}

// Attributes nolint
type Attributes []Attribute

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

// Amounts - the array of Amount
type Amounts []Amount

func (amounts Amounts) toCurrencies() (currenies []blockatlas.Currency) {
	for _, amt := range amounts {
		currenies = append(currenies, blockatlas.Currency{
			Decimals:   coin.Terra().Decimals,
			Symbol:     DenomMap[amt.Denom],
			Value:      blockatlas.Amount(amt.Quantity),
			CurrencyID: amt.Denom,
		})
	}
	return
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
	case MsgSend:
		var msgTransfer MessageValueTransfer
		err = json.Unmarshal(messageInternal.Value, &msgTransfer)
		m.Value = msgTransfer
	}
	return err
}

// AuthAccount is response body of account query
type AuthAccount struct {
	Account Account `json:"result"`
}

// Account nolint
type Account struct {
	Value AccountValue `json:"value"`
}

// AccountValue nolint
type AccountValue struct {
	Coins []Balance `json:"coins"`
}

// Balance nolint
type Balance struct {
	Denom  DenomType `json:"denom"`
	Amount string    `json:"amount"`
}

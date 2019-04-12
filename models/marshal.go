package models

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/trustwallet/blockatlas/util"
	"regexp"
	"sort"
	"strings"
)

var matchNumber = regexp.MustCompile(`^\d+(\.\d+)?$`)

// Tx, but with default JSON marshalling methods
type wrappedTx Tx

func (t *Tx) UnmarshalJSON(data []byte) error {
	// Wrap the Tx type to avoid infinite recursion
	var wrapped wrappedTx

	var raw json.RawMessage
	wrapped.Meta = &raw
	if err := json.Unmarshal(data, &wrapped); err != nil {
		return err
	}

	*t = Tx(wrapped)

	switch t.Type {
	case TxTransfer:
		t.Meta = new(Transfer)
	case TxTokenTransfer:
		t.Meta = new(TokenTransfer)
	case TxCollectibleTransfer:
		t.Meta = new(CollectibleTransfer)
	case TxTokenSwap:
		t.Meta = new(TokenSwap)
	case TxContractCall:
		t.Meta = new(ContractCall)
	default:
		return fmt.Errorf(`unsupported tx type "%s"`, t.Type)
	}
	if err := json.Unmarshal(raw, t.Meta); err != nil {
		return err
	}
	return nil
}

func (t *Tx) MarshalJSON() ([]byte, error) {
	// Set type from metadata content
	switch t.Meta.(type) {
	case Transfer, *Transfer:
		t.Type = TxTransfer
	case TokenTransfer, *TokenTransfer:
		t.Type = TxTokenTransfer
	case CollectibleTransfer, *CollectibleTransfer:
		t.Type = TxCollectibleTransfer
	case TokenSwap, *TokenSwap:
		t.Type = TxTokenSwap
	case ContractCall, *ContractCall:
		t.Type = TxContractCall
	default:
		return nil, fmt.Errorf("unsupported tx metadata")
	}

	// Set status to completed by default
	if t.Status == "" {
		t.Status = StatusCompleted
	}

	// Wrap the Tx type to avoid infinite recursion
	return json.Marshal(wrappedTx(*t))
}

func (a *Amount) UnmarshalJSON(data []byte) error {
	var n json.Number
	err := json.Unmarshal(data, &n)
	if err != nil {
		return err
	}
	str := string(n)
	if !matchNumber.MatchString(str) {
		return fmt.Errorf("not a regular decimal number")
	}
	if strings.ContainsRune(str, '.') {
		str, _ = util.DecimalToSatoshis(str)
	}
	*a = Amount(str)
	return nil
}

func (a *Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*a))
}

func (r *Response) Sort() {
	sort.Slice(*r, func(i, j int) bool {
		ti := cast.ToUint64((*r)[i].Date)
		tj := cast.ToUint64((*r)[j].Date)
		return ti >= tj
	})
}

func (r *Response) MarshalJSON() ([]byte, error) {
	var page struct {
		Total  int    `json:"total"`
		Docs   []Tx   `json:"docs"`
		Status string `json:"status"`
	}
	page.Docs = []Tx(*r)
	if page.Docs == nil {
		page.Docs = make([]Tx, 0)
	}
	page.Total = len(page.Docs)
	page.Status = "success"
	return json.Marshal(page)
}

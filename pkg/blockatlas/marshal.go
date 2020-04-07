package blockatlas

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"regexp"
	"strings"
)

var matchNumber = regexp.MustCompile(`^\d+(\.\d+)?$`)

// Tx, but with default JSON marshalling methods
type wrappedTx Tx

// UnmarshalJSON creates a transaction along with metadata from a JSON object.
// Fails if the meta object can't be read.
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
	case TxMultiCurrencyTransfer:
		t.Meta = new(MultiCurrencyTransfer)
	case TxNativeTokenTransfer:
		t.Meta = new(NativeTokenTransfer)
	case TxTokenTransfer:
		t.Meta = new(TokenTransfer)
	case TxCollectibleTransfer:
		t.Meta = new(CollectibleTransfer)
	case TxTokenSwap:
		t.Meta = new(TokenSwap)
	case TxContractCall:
		t.Meta = new(ContractCall)
	case TxAnyAction:
		t.Meta = new(AnyAction)
	default:
		return errors.E("unsupported tx type", errors.Params{"type": t.Type})
	}
	if err := json.Unmarshal(raw, t.Meta); err != nil {
		return err
	}
	return nil
}

// MarshalJSON creates a JSON object from a transaction.
// Sets the Type field to the currect value based on the Meta type.
func (t *Tx) MarshalJSON() ([]byte, error) {
	// Set type from metadata content
	switch t.Meta.(type) {
	case Transfer, *Transfer:
		t.Type = TxTransfer
	case MultiCurrencyTransfer, *MultiCurrencyTransfer:
		t.Type = TxMultiCurrencyTransfer
	case NativeTokenTransfer, *NativeTokenTransfer:
		t.Type = TxNativeTokenTransfer
	case TokenTransfer, *TokenTransfer:
		t.Type = TxTokenTransfer
	case CollectibleTransfer, *CollectibleTransfer:
		t.Type = TxCollectibleTransfer
	case TokenSwap, *TokenSwap:
		t.Type = TxTokenSwap
	case ContractCall, *ContractCall:
		t.Type = TxContractCall
	case AnyAction, *AnyAction:
		t.Type = TxAnyAction
	default:
		return nil, errors.E("unsupported tx metadata", errors.Params{"meta": t.Meta})
	}

	// Set status to completed by default
	if t.Status == "" {
		t.Status = StatusCompleted
	}

	// Wrap the Tx type to avoid infinite recursion
	return json.Marshal(wrappedTx(*t))
}

// UnmarshalJSON reads an amount from a JSON string or number.
// Comma separators get dropped with address.DecimalToSatoshis.
func (a *Amount) UnmarshalJSON(data []byte) error {
	var n json.Number
	err := json.Unmarshal(data, &n)
	if err != nil {
		return err
	}
	str := string(n)
	if !matchNumber.MatchString(str) {
		return errors.E("not a regular decimal number", errors.Params{"str": str})
	}
	if strings.ContainsRune(str, '.') {
		str, _ = numbers.DecimalToSatoshis(str)
	}
	*a = Amount(str)
	return nil
}

// MarshalJSON returns a JSON string representing the amount
func (a *Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*a))
}

// Sort sorts the response by date, descending
func (txs TxPage) Len() int           { return len(txs) }
func (txs TxPage) Less(i, j int) bool { return txs[i].Date > txs[j].Date }
func (txs TxPage) Swap(i, j int)      { txs[i], txs[j] = txs[j], txs[i] }

// MarshalJSON returns a wrapped list of transactions in JSON
func (r *TxPage) MarshalJSON() ([]byte, error) {
	var page struct {
		Total  int  `json:"total"`
		Docs   []Tx `json:"docs"`
		Status bool `json:"status"`
	}
	page.Docs = *r
	if page.Docs == nil {
		page.Docs = make([]Tx, 0)
	}
	page.Total = len(page.Docs)
	page.Status = true
	return json.Marshal(page)
}

// MarshalJSON returns a wrapped list of collections in JSON
func (r CollectionPage) MarshalJSON() ([]byte, error) {
	var page struct {
		Total  int          `json:"total"`
		Docs   []Collection `json:"docs"`
		Status bool         `json:"status"`
	}
	page.Docs = []Collection(r)
	if page.Docs == nil {
		page.Docs = make([]Collection, 0)
	}
	page.Total = len(page.Docs)
	page.Status = true
	return json.Marshal(page)
}

// MarshalJSON returns a wrapped list of collectibles in JSON
func (r CollectiblePage) MarshalJSON() ([]byte, error) {
	var page struct {
		Total  int           `json:"total"`
		Docs   []Collectible `json:"docs"`
		Status bool          `json:"status"`
	}
	page.Docs = []Collectible(r)
	if page.Docs == nil {
		page.Docs = make([]Collectible, 0)
	}
	page.Total = len(page.Docs)
	page.Status = true
	return json.Marshal(page)
}

// MarshalJSON returns a wrapped list of collections in JSON
func (r CollectionPageV3) MarshalJSON() ([]byte, error) {
	var page struct {
		Total  int            `json:"total"`
		Docs   []CollectionV3 `json:"docs"`
		Status bool           `json:"status"`
	}
	page.Docs = []CollectionV3(r)
	if page.Docs == nil {
		page.Docs = make([]CollectionV3, 0)
	}
	page.Total = len(page.Docs)
	page.Status = true
	return json.Marshal(page)
}

// MarshalJSON returns a wrapped list of collectibles in JSON
func (r CollectiblePageV3) MarshalJSON() ([]byte, error) {
	var page struct {
		Total  int             `json:"total"`
		Docs   []CollectibleV3 `json:"docs"`
		Status bool            `json:"status"`
	}
	page.Docs = []CollectibleV3(r)
	if page.Docs == nil {
		page.Docs = make([]CollectibleV3, 0)
	}
	page.Total = len(page.Docs)
	page.Status = true
	return json.Marshal(page)
}

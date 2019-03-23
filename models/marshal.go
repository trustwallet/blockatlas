package models

import (
	"encoding/json"
	"fmt"
)

func (t *Tx) UnmarshalJSON(data []byte) error {
	// Wrap the Tx type to avoid infinite
	// recursion in UnmarshalJSON
	type wrappedTx Tx
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

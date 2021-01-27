package solana

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

const (
	TypeTransfer = "transfer"
)

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	results := make(types.TxPage, 0)
	txs, err := p.client.GetTransactions(address)
	if err != nil {
		return results, err
	}
	for _, tx := range txs {
		if normalized, err := NormalizeTx(tx, address); err == nil {
			results = append(results, normalized)
		}
	}
	return results, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (types.TxPage, error) {
	results := make(types.TxPage, 0)
	value, err := p.client.GetTokenAccountsByOwner(address, token)
	if err != nil {
		return results, err
	}
	if len(value.Value) == 0 {
		return results, fmt.Errorf("not token (%s) account for %s", token, address)
	}
	pubkey := value.Value[0].Pubkey
	txs, err := p.client.GetTransactions(pubkey)

	for _, tx := range txs {
		if normalized, err := NormalizeTokenTx(tx, address, pubkey, token); err == nil {
			results = append(results, normalized)
		}
	}
	return results, nil
}

func ensureInstruction(tx ConfirmedTransaction) error {
	// only check first instruction
	if len(tx.Transaction.Message.Instructions) != 1 || len(tx.Transaction.Signatures) != 1 {
		return errors.New("not supported")
	}
	// only supports transfer type now
	instruction := tx.Transaction.Message.Instructions[0]
	if instruction.Parsed.Type != TypeTransfer {
		return errors.New("not supported type other than transfer")
	}
	return nil
}

func getTransactionStatus(tx ConfirmedTransaction) types.Status {
	// tx status
	status := types.StatusCompleted
	if tx.Meta.Err != nil {
		status = types.StatusError
	}
	return status
}

func NormalizeTx(tx ConfirmedTransaction, address string) (normalized types.Tx, err error) {
	err = ensureInstruction(tx)
	if err != nil {
		return
	}

	instruction := tx.Transaction.Message.Instructions[0]
	from := instruction.Parsed.Info.Source

	// tx direction
	direction := types.DirectionIncoming
	if address == from {
		direction = types.DirectionOutgoing
	}

	coin := coin.Solana()
	normalized = types.Tx{
		ID:        tx.Transaction.Signatures[0],
		Coin:      coin.ID,
		From:      from,
		To:        instruction.Parsed.Info.Destination,
		Fee:       types.Amount(strconv.FormatUint(tx.Meta.Fee, 10)),
		Date:      EstimateTimestamp(tx.Slot),
		Block:     tx.Slot,
		Status:    getTransactionStatus(tx),
		Type:      types.TxTransfer,
		Direction: direction,
		Meta: types.Transfer{
			Value:    types.Amount(strconv.FormatUint(instruction.Parsed.Info.Lamports, 10)),
			Symbol:   coin.Symbol,
			Decimals: coin.Decimals,
		},
	}

	return normalized, nil
}

func NormalizeTokenTx(tx ConfirmedTransaction, address, pubkey, token string) (normalized types.Tx, err error) {
	err = ensureInstruction(tx)
	if err != nil {
		return
	}

	instruction := tx.Transaction.Message.Instructions[0]
	from := instruction.Parsed.Info.Source

	direction := types.DirectionIncoming
	if pubkey == from {
		direction = types.DirectionOutgoing
	}

	// Metadata is not set, hold for v2 transaction format
	normalized = types.Tx{
		ID:        tx.Transaction.Signatures[0],
		Coin:      coin.SOLANA,
		From:      from,
		To:        instruction.Parsed.Info.Destination,
		Fee:       types.Amount(strconv.FormatUint(tx.Meta.Fee, 10)),
		Date:      EstimateTimestamp(tx.Slot),
		Block:     tx.Slot,
		Status:    getTransactionStatus(tx),
		Type:      types.TxTokenTransfer,
		Direction: direction,
	}

	return
}

func EstimateTimestamp(slot uint64) int64 {
	var (
		blockTime  uint64 = 400 //ms
		sampleSlot uint64 = 52838300
		sampleTs   uint64 = 1606944859 * 1000
	)
	offset := (slot - sampleSlot) * blockTime
	return int64((sampleTs + offset) / 1000)
}

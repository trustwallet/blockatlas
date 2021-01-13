package solana

import (
	"errors"
	"strconv"

	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	results := make(types.TxPage, 0)
	txs, err := p.client.GetTransactions(address)
	if err != nil {
		return results, err
	}
	for _, tx := range txs {
		if normalized, err := p.NormalizeTx(tx, address); err == nil {
			results = append(results, normalized)
		}
	}
	return results, nil
}

func (p *Platform) NormalizeTx(tx ConfirmedTransaction, address string) (normalized types.Tx, err error) {

	// only check first instruction
	if len(tx.Transaction.Message.Instructions) != 1 || len(tx.Transaction.Signatures) != 1 {
		return normalized, errors.New("not supported")
	}

	// only supports transfer type now
	instruction := tx.Transaction.Message.Instructions[0]
	if instruction.Parsed.Type != "transfer" {
		return normalized, errors.New("not supported type other than transfer")
	}

	// tx direction
	from := instruction.Parsed.Info.Source
	direction := types.DirectionIncoming
	if address == from {
		direction = types.DirectionOutgoing
	}

	// tx status
	status := types.StatusCompleted
	if tx.Meta.Err != nil {
		status = types.StatusError
	}

	normalized = types.Tx{
		ID:        tx.Transaction.Signatures[0],
		Coin:      p.Coin().ID,
		From:      from,
		To:        instruction.Parsed.Info.Destination,
		Fee:       types.Amount(strconv.FormatUint(tx.Meta.Fee, 10)),
		Date:      EstimateTimestamp(tx.Slot),
		Block:     tx.Slot,
		Status:    status,
		Type:      types.TxTransfer,
		Direction: direction,
		Meta: types.Transfer{
			Value:    types.Amount(strconv.FormatUint(instruction.Parsed.Info.Lamports, 10)),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}

	return normalized, nil
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

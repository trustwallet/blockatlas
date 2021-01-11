package solana

import (
	"errors"
	"strconv"

	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	results := make(txtype.TxPage, 0)
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

func (p *Platform) NormalizeTx(tx ConfirmedTransaction, address string) (normalized txtype.Tx, err error) {

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
	direction := txtype.DirectionIncoming
	if address == from {
		direction = txtype.DirectionOutgoing
	}

	// tx status
	status := txtype.StatusCompleted
	if tx.Meta.Err != nil {
		status = txtype.StatusError
	}

	normalized = txtype.Tx{
		ID:        tx.Transaction.Signatures[0],
		Coin:      p.Coin().ID,
		From:      from,
		To:        instruction.Parsed.Info.Destination,
		Fee:       txtype.Amount(strconv.FormatUint(tx.Meta.Fee, 10)),
		Date:      EstimateTimestamp(tx.Slot),
		Block:     tx.Slot,
		Status:    status,
		Type:      txtype.TxTransfer,
		Direction: direction,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(strconv.FormatUint(instruction.Parsed.Info.Lamports, 10)),
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

package solana

import (
	"errors"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	results := make(blockatlas.TxPage, 0)
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

func (p *Platform) NormalizeTx(tx ConfirmedTransaction, address string) (normalized blockatlas.Tx, err error) {

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
	direction := blockatlas.DirectionIncoming
	if address == from {
		direction = blockatlas.DirectionOutgoing
	}

	// tx status
	status := blockatlas.StatusCompleted
	if tx.Meta.Err != nil {
		status = blockatlas.StatusError
	}

	normalized = blockatlas.Tx{
		ID:        tx.Transaction.Signatures[0],
		Coin:      p.Coin().ID,
		From:      from,
		To:        instruction.Parsed.Info.Destination,
		Fee:       blockatlas.Amount(strconv.FormatUint(tx.Meta.Fee, 10)),
		Block:     tx.Slot,
		Status:    status,
		Type:      blockatlas.TxTransfer,
		Direction: direction,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(strconv.FormatUint(instruction.Parsed.Info.Lamports, 10)),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}

	return normalized, nil
}

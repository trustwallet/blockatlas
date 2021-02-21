package solana

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/types"
)

const (
	programSystem = "system"
	programToken  = "spl-token"

	instructionTransfer = "transfer"
	// will support instructionTransferChecked later
	// instructionTransferChecked = "transferChecked"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	results := make(types.Txs, 0)
	txs, err := p.client.GetTransactionsByAddress(address)
	if err != nil {
		return results, err
	}
	for _, tx := range txs {
		if tx.BlockTime == 0 {
			continue
		}
		if normalized, err := p.NormalizeTx(tx, tx.Slot, tx.BlockTime); err == nil {
			normalized.Direction = normalized.GetTransactionDirection(address)
			results = append(results, normalized)
		}
	}
	return results, nil
}

func (p *Platform) NormalizeTx(tx ConfirmedTransaction, slot uint64, timestamp int64) (normalized types.Tx, err error) {

	// only check first instruction
	if len(tx.Transaction.Message.Instructions) != 1 || len(tx.Transaction.Signatures) != 1 {
		err = errors.New("not supported instructions/signatures count")
		return
	}

	// only supports transfer and token transfer
	instruction := tx.Transaction.Message.Instructions[0]

	if instruction.Program != programSystem && instruction.Program != programToken {
		err = fmt.Errorf("not supported program: %s", instruction.Program)
		return
	}

	var parsed Parsed
	err = blockatlas.MapJsonObject(instruction.Parsed, &parsed)

	if err != nil {
		return
	}

	// tx status
	status := types.StatusCompleted
	if tx.Meta.Err != nil {
		status = types.StatusError
	}

	normalized = types.Tx{
		ID:     tx.Transaction.Signatures[0],
		Coin:   p.Coin().ID,
		Fee:    types.Amount(strconv.FormatUint(tx.Meta.Fee, 10)),
		Date:   timestamp,
		Block:  slot,
		Status: status,
	}

	switch parsed.Type {
	case instructionTransfer:
		var transfer TransferInfo
		err = blockatlas.MapJsonObject(parsed.Info, &transfer)
		if err == nil {
			normalized.From = transfer.Source
			normalized.To = transfer.Destination
			normalized.Type = types.TxTransfer
			normalized.Meta = types.Transfer{
				Value:    types.Amount(strconv.FormatUint(transfer.Lamports, 10)),
				Symbol:   p.Coin().Symbol,
				Decimals: p.Coin().Decimals,
			}
		}
	default:
		err = fmt.Errorf("not supported type: %s", parsed.Type)
	}
	return
}

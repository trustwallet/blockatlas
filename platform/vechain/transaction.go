package vechain

import (
	"errors"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/address"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTokenTxsByAddress(address, token string) (page types.Txs, err error) {
	if token != gasTokenAddress {
		return
	}

	blockNumber, err := p.CurrentBlockNumber()
	if err != nil {
		return
	}

	events, err := p.client.GetLogsEvent(address, token, blockNumber)
	if err != nil {
		return
	}

	eventsIDs := make([]string, 0)
	for _, event := range events {
		eventsIDs = append(eventsIDs, event.Meta.TxId)
	}

	txs, err := p.getTransactionsByIDs(eventsIDs)
	if err != nil {
		return
	}

	// NormalizeTokenTransaction won't set tx direction anymore, set it here
	for _, tx := range txs {
		updateTransactionDirection(&tx, address)
	}

	return txs, nil
}

func (p *Platform) getTransactionsByIDs(ids []string) (types.Txs, error) {
	page := types.Txs{}
	for _, id := range ids {
		tx, err := p.client.GetTransactionByID(id)
		if err != nil {
			return nil, err
		}

		receipt, err := p.client.GetTransactionReceiptByID(id)
		if err != nil {
			return page, err
		}

		txs, err := NormalizeTokenTransaction(tx, receipt)
		if err != nil {
			return page, err
		}
		page = append(page, txs...)
	}
	return page, nil
}

func NormalizeTokenTransaction(srcTx Tx, receipt TxReceipt) (types.Txs, error) {
	// the only supported Token on VeChain is its Gas token
	if receipt.Reverted {
		return normalizeRevertedTokenTransaction(srcTx, receipt)
	}

	txs := make(types.Txs, 0)

	if receipt.Outputs == nil || len(receipt.Outputs) == 0 {
		return types.Txs{}, errors.New("NormalizeBlockTransaction: receipt.Outputs not found: " + srcTx.Id)
	}

	for _, output := range receipt.Outputs {
		if len(output.Events) == 0 || len(output.Events[0].Topics) < 3 {
			continue
		}

		fee, err := numbers.HexToDecimal(receipt.Paid)
		if err != nil {
			return txs, err
		}

		originSender, err := address.EIP55Checksum(blockatlas.GetValidParameter(srcTx.Origin, srcTx.Meta.TxOrigin))
		if err != nil {
			return txs, err
		}

		event := output.Events[0]

		value, err := numbers.HexToDecimal(event.Data)
		if err != nil {
			continue
		}

		originReceiver, err := address.EIP55Checksum(event.Address)
		if err != nil {
			continue
		}

		if originReceiver != gasTokenAddress {
			continue
		}

		topicsTo, err := address.EIP55Checksum(getRecipientAddress(event.Topics[2]))
		if err != nil {
			continue
		}

		txs = append(txs, types.Tx{
			ID:     srcTx.Id,
			Coin:   coin.VECHAIN,
			From:   originSender,
			To:     originReceiver,
			Fee:    types.Amount(fee),
			Date:   srcTx.Meta.BlockTimestamp,
			Type:   types.TxTokenTransfer,
			Block:  srcTx.Meta.BlockNumber,
			Status: types.StatusCompleted,
			Meta: types.TokenTransfer{
				Name:     gasTokenName,
				TokenID:  originReceiver,
				Value:    types.Amount(value),
				Symbol:   gasTokenSymbol,
				Decimals: gasTokenDecimals,
				From:     originSender,
				To:       topicsTo,
			},
		})
	}
	return txs, nil
}

func normalizeRevertedTokenTransaction(srcTx Tx, receipt TxReceipt) (types.Txs, error) {
	txs := make(types.Txs, 0)

	fee, err := numbers.HexToDecimal(receipt.Paid)
	if err != nil {
		return txs, err
	}

	originSender, err := address.EIP55Checksum(blockatlas.GetValidParameter(srcTx.Origin, srcTx.Meta.TxOrigin))
	if err != nil {
		return txs, err
	}

	var to string
	if len(srcTx.Clauses) > 0 {
		to = srcTx.Clauses[0].To
		if checksumTo, err := address.EIP55Checksum(to); err == nil {
			to = checksumTo
		}
	} else {
		return txs, errors.New("NormalizeBlockTransaction: srcTx.Clauses not found: " + srcTx.Id)
	}

	if to != gasTokenAddress {
		return txs, nil
	}

	txs = append(txs, types.Tx{
		ID:     srcTx.Id,
		Coin:   coin.VECHAIN,
		From:   originSender,
		To:     to,
		Fee:    types.Amount(fee),
		Date:   srcTx.Meta.BlockTimestamp,
		Type:   types.TxTokenTransfer,
		Block:  srcTx.Meta.BlockNumber,
		Status: types.StatusError,
		Meta: types.TokenTransfer{
			Name:     gasTokenName,
			TokenID:  gasTokenAddress,
			Value:    "0",
			Symbol:   gasTokenSymbol,
			Decimals: gasTokenDecimals,
			From:     originSender,
			To:       to,
		},
	})
	return txs, nil
}

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	headBlock, err := p.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}
	transfers, err := p.client.GetTransactions(address, headBlock)
	if err != nil {
		return nil, err
	}

	txs := make(types.Txs, 0)
	for _, t := range transfers {
		trxId, err := p.client.GetTransactionByID(t.Meta.TxId)
		if err != nil {
			continue
		}
		tx, err := p.NormalizeTransaction(t, trxId, address)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (p *Platform) NormalizeTransaction(srcTx LogTransfer, trxId Tx, addr string) (tx types.Tx, err error) {
	value, err := numbers.HexToDecimal(srcTx.Amount)
	if err != nil {
		return
	}

	fee := strconv.Itoa(trxId.Gas)
	sender, err := address.EIP55Checksum(srcTx.Sender)
	if err != nil {
		return
	}
	recipient, err := address.EIP55Checksum(srcTx.Recipient)
	if err != nil {
		return
	}
	addrChecksum, err := address.EIP55Checksum(addr)
	if err != nil {
		return
	}

	direction, err := getTransactionDirection(sender, recipient, addrChecksum)
	if err != nil {
		return types.Tx{}, err
	}

	return types.Tx{
		ID:        srcTx.Meta.TxId,
		Coin:      p.Coin().ID,
		From:      sender,
		To:        recipient,
		Fee:       types.Amount(fee),
		Date:      srcTx.Meta.BlockTimestamp,
		Type:      types.TxTransfer,
		Block:     srcTx.Meta.BlockNumber,
		Direction: direction,
		Status:    types.StatusCompleted,
		Meta: types.Transfer{
			Value:    types.Amount(value),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}, nil
}

func hexToInt(hex string) (uint64, error) {
	nonceStr, err := numbers.HexToDecimal(hex)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(nonceStr, 10, 64)
}

// Substring recipient address from clause data and appends 0x
// 0x000000000000000000000000b5e883349e68ab59307d1604555ac890fac47128 => 0xb5e883349e68ab59307d1604555ac890fac47128
func getRecipientAddress(hex string) string {
	return "0x" + hex[len(hex)-40:]
}

func updateTransactionDirection(tx *types.Tx, addr string) {
	meta, ok := tx.Meta.(types.TokenTransfer)
	if !ok {
		return
	}
	direction, err := getTransactionDirection(tx.From, meta.To, addr)
	if err != nil {
		return
	}
	tx.Direction = direction
}

func getTransactionDirection(sender, recipient, addr string) (types.Direction, error) {
	if sender != addr && recipient != addr {
		return "", errors.New("Unknown direction")
	}
	if sender == addr && recipient == addr {
		return types.DirectionSelf, nil
	}
	if sender == addr {
		return types.DirectionOutgoing, nil
	} else {
		return types.DirectionIncoming, nil
	}
}

package vechain

import (
	"errors"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/address"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTokenTxsByAddress(address, token string) (types.TxPage, error) {
	if token != gasTokenAddress {
		return nil, nil
	}
	curBlock, err := p.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}
	events, err := p.client.GetLogsEvent(address, token, curBlock)
	if err != nil {
		return nil, err
	}
	eventsIDs := make([]string, 0)
	for _, event := range events {
		eventsIDs = append(eventsIDs, event.Meta.TxId)
	}

	cTxs := p.getTransactionsByIDs(eventsIDs)
	txs := make(types.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return txs, nil
}

func (p *Platform) getTransactionsByIDs(ids []string) chan types.TxPage {
	txChan := make(chan types.TxPage, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, c chan types.TxPage) {
			defer wg.Done()
			err := p.getTransactionChannel(i, c)
			if err != nil {
				log.Error(err)
			}
		}(id, txChan)
	}
	wg.Wait()
	close(txChan)
	return txChan
}

func (p *Platform) getTransactionChannel(id string, txChan chan types.TxPage) error {
	srcTx, err := p.client.GetTransactionByID(id)
	if err != nil {
		return err
	}

	receipt, err := p.client.GetTransactionReceiptByID(id)
	if err != nil {
		return err
	}

	txs, err := p.NormalizeTokenTransaction(srcTx, receipt)
	if err != nil {
		return err
	}
	txChan <- txs
	return nil
}

func (p *Platform) NormalizeTokenTransaction(srcTx Tx, receipt TxReceipt) (types.TxPage, error) {
	txs := make(types.TxPage, 0)

	fee, err := numbers.HexToDecimal(receipt.Paid)
	if err != nil {
		return txs, err
	}

	originSender, err := address.EIP55Checksum(blockatlas.GetValidParameter(srcTx.Origin, srcTx.Meta.TxOrigin))
	if err != nil {
		return txs, err
	}

	if receipt.Reverted {
		var to string
		if len(srcTx.Clauses) > 0 {
			to = srcTx.Clauses[0].To
			if checksumTo, err := address.EIP55Checksum(to); err == nil {
				to = checksumTo
			}
		} else {
			return txs, errors.New("NormalizeBlockTransaction: srcTx.Clauses not found: " + srcTx.Id)
		}

		txs = append(txs, types.Tx{
			ID:     srcTx.Id,
			Coin:   p.Coin().ID,
			From:   originSender,
			To:     to,
			Fee:    types.Amount(fee),
			Date:   srcTx.Meta.BlockTimestamp,
			Type:   types.TxTokenTransfer,
			Block:  srcTx.Meta.BlockNumber,
			Status: types.StatusError,
		})
		return txs, nil
	}

	if receipt.Outputs == nil || len(receipt.Outputs) == 0 {
		return types.TxPage{}, errors.New("NormalizeBlockTransaction: receipt.Outputs not found: " + srcTx.Id)
	}

	for _, output := range receipt.Outputs {
		if len(output.Events) == 0 || len(output.Events[0].Topics) < 3 {
			continue
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
		topicsFrom, err := address.EIP55Checksum(getRecipientAddress(event.Topics[1]))
		if err != nil {
			continue
		}
		topicsTo, err := address.EIP55Checksum(getRecipientAddress(event.Topics[2]))
		if err != nil {
			continue
		}

		direction, err := getTokenTransactionDirectory(originSender, topicsFrom, topicsTo)
		if err != nil {
			continue
		}

		txs = append(txs, types.Tx{
			ID:        srcTx.Id,
			Coin:      p.Coin().ID,
			From:      originSender,
			To:        originReceiver,
			Fee:       types.Amount(fee),
			Date:      srcTx.Meta.BlockTimestamp,
			Type:      types.TxTokenTransfer,
			Block:     srcTx.Meta.BlockNumber,
			Status:    types.StatusCompleted,
			Direction: direction,
			Meta: types.TokenTransfer{
				// the only supported Token on VeChain is its Gas token
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

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	headBlock, err := p.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}
	transfers, err := p.client.GetTransactions(address, headBlock)
	if err != nil {
		return nil, err
	}

	txs := make(types.TxPage, 0)
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

	directory, err := getTransferDirectory(sender, recipient, addrChecksum)
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
		Direction: directory,
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

func getTokenTransactionDirectory(originSender, topicsFrom, topicsTo string) (dir types.Direction, err error) {
	if originSender == topicsFrom && originSender == topicsTo {
		return types.DirectionSelf, nil
	}
	if originSender == topicsFrom && originSender != topicsTo {
		return types.DirectionIncoming, nil
	}
	if originSender == topicsTo && originSender != topicsFrom {
		return types.DirectionOutgoing, nil
	}
	return "", errors.New("Unknown direction")
}

func getTransferDirectory(sender, recipient, addr string) (dir types.Direction, err error) {
	if sender == addr && recipient == addr {
		return types.DirectionSelf, nil
	}
	if sender == addr && recipient != addr {
		return types.DirectionOutgoing, nil
	}
	if recipient == addr && sender != addr {
		return types.DirectionIncoming, nil
	}
	return "", errors.New("Unknown direction")
}

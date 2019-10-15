package vechain

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/util"
	"strings"
	"sync"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("vechain.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.VET]
}

const (
	GasContract         = "0x0000000000000000000000000000456e65726779"
	GasSymbol           = "VTHO"
	GasName             = "VeThor Token"
	VeThorTransferEvent = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

// CurrentBlockNumber implementation of interface function which gets a current blockchain height
func (p *Platform) CurrentBlockNumber() (int64, error) {
	cbi, err := p.client.GetCurrentBlockInfo()
	if err != nil {
		return 0, err
	}
	return cbi.BestBlockNum, nil
}

// GetBlockByNumber implementation of interface function which gets a block for push notification
func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	transactionsChan := p.getTransactions(block.Transactions)

	txs := make([]blockatlas.Tx, 0)
	for t := range transactionsChan {
		ntxs := NormalizeBlockTransactions(t)
		txs = append(txs, ntxs...)
	}

	return &blockatlas.Block{
		Number: num,
		ID:     block.ID,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.getTxsByAddress(address)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	if strings.ToLower(token) == GasContract {
		return p.getThorTxsByAddress(address)
	} else {
		return nil, errors.E("vechain invalid token", errors.TypePlatformUnmarshal,
			errors.Params{"token": token}).PushToSentry()
	}
}

func (p *Platform) getThorTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, err := p.client.GetTokenTransfers(address)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, tx := range sourceTxs.TokenTransfers {
		ids = append(ids, tx.TxID)
	}
	receiptsChan := p.getTransactionReceipt(ids)

	var txs []blockatlas.Tx
	for _, t := range sourceTxs.TokenTransfers {
		if !strings.EqualFold(t.ContractAddress, GasContract) {
			continue
		}

		receipt := findTransferReceiptByTxID(receiptsChan, t.TxID)
		if receipt == nil {
			logger.Error("findTransferReceiptByTxID cannot find the receipt", logger.Params{
				"transfer": t,
			})
			continue
		}

		tx, err := NormalizeTokenTransfer(&t, receipt)
		if err != nil {
			p := logger.Params{"receipt": receipt, "transfer": t}
			err = errors.E(err, "invalid token", errors.TypePlatformUnmarshal, p).PushToSentry()
			logger.Error(err, "getTxsByAddress clause error", p)
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func (p *Platform) getTransactionReceipt(ids []string) chan *TransferReceipt {
	receiptsChan := make(chan *TransferReceipt, len(ids))
	sem := util.NewSemaphore(16)
	var wg sync.WaitGroup
	wg.Add(len(ids))
	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			receipt, err := p.client.GetTransactionReceipt(id)
			if err != nil {
				err = errors.E(err, "Failed to get tx receipt", errors.TypePlatformUnmarshal,
					errors.Params{"id": id}).PushToSentry()
				logger.Error(err, logger.Params{"id": id})
				return
			}
			receiptsChan <- receipt
		}(id)
	}
	wg.Wait()
	close(receiptsChan)

	return receiptsChan
}

func (p *Platform) getTransactions(ids []string) chan *NativeTransaction {
	receiptsChan := make(chan *NativeTransaction, len(ids))

	sem := util.NewSemaphore(16)
	var wg sync.WaitGroup
	wg.Add(len(ids))
	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			receipt, err := p.client.GetTransactionByID(id)
			if err != nil {
				err = errors.E(err, "Failed to get tx transaction", errors.TypePlatformUnmarshal,
					errors.Params{"id": id}).PushToSentry()
				logger.Error(err, logger.Params{"id": id})
				return
			}
			receiptsChan <- receipt
		}(id)
	}

	wg.Wait()
	close(receiptsChan)

	return receiptsChan
}

func findTransferReceiptByTxID(receiptsChan chan *TransferReceipt, txID string) *TransferReceipt {
	for receipt := range receiptsChan {
		if receipt.ID == txID {
			return receipt
		}
	}
	return nil
}

func (p *Platform) getTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTransactions(address)

	var ids []string
	for _, tx := range sourceTxs.Transactions {
		ids = append(ids, tx.ID)
	}
	receiptsChan := p.getTransactionReceipt(ids)

	var txs []blockatlas.Tx
	for receipt := range receiptsChan {
		for _, clause := range receipt.Clauses {
			if !strings.EqualFold(receipt.Origin, address) && !strings.EqualFold(clause.To, address) {
				continue
			}
			tx, err := NormalizeTransfer(receipt, &clause)
			if err != nil {
				p := logger.Params{"receipt": receipt, "clause": clause}
				err = errors.E(err, errors.TypePlatformUnmarshal, p).PushToSentry()
				logger.Error(err, "getTxsByAddress clause error", p)
				continue
			}
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func NormalizeTransfer(receipt *TransferReceipt, clause *Clause) (blockatlas.Tx, error) {
	if receipt.Receipt == nil || clause == nil {
		return blockatlas.Tx{}, errors.E("invalid parameters", errors.Params{"receipt": receipt, "clause": clause}).PushToSentry()
	}
	feeBase10, err := util.HexToDecimal(receipt.Receipt.Paid)
	if err != nil {
		return blockatlas.Tx{}, err
	}
	valueBase10, err := util.HexToDecimal(clause.Value)
	if err != nil {
		return blockatlas.Tx{}, err
	}
	tx := blockatlas.Tx{
		ID:       receipt.ID,
		Coin:     coin.VET,
		From:     receipt.Origin,
		To:       clause.To,
		Fee:      blockatlas.Amount(feeBase10),
		Date:     int64(receipt.Timestamp),
		Type:     blockatlas.TxTransfer,
		Block:    receipt.Block,
		Status:   ReceiptStatus(receipt.Receipt.Reverted),
		Sequence: receipt.Block,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(valueBase10),
			Symbol:   coin.Coins[coin.VET].Symbol,
			Decimals: 18,
		},
	}
	return tx, nil
}

func NormalizeTokenTransfer(t *TokenTransfer, receipt *TransferReceipt) (blockatlas.Tx, error) {
	feeBase10, err := util.HexToDecimal(receipt.Receipt.Paid)
	if err != nil {
		return blockatlas.Tx{}, err
	}
	valueBase10, err := util.HexToDecimal(t.Amount)
	if err != nil {
		return blockatlas.Tx{}, err
	}

	tx := blockatlas.Tx{
		ID:       t.TxID,
		Coin:     coin.VET,
		From:     t.Origin,
		To:       t.Receiver,
		Fee:      blockatlas.Amount(feeBase10),
		Date:     t.Timestamp,
		Type:     blockatlas.TxNativeTokenTransfer,
		Block:    t.Block,
		Status:   ReceiptStatus(receipt.Receipt.Reverted),
		Sequence: t.Block,
		Meta: blockatlas.NativeTokenTransfer{
			Name:     GasName,
			Symbol:   GasSymbol,
			TokenID:  GasContract,
			Decimals: 18,
			Value:    blockatlas.Amount(valueBase10),
			From:     t.Origin,
			To:       t.Receiver,
		},
	}
	return tx, nil
}

// NormalizeTransaction converts a VeChain VTHO token transaction into the generic model
func NormalizeBlockTransactions(t *NativeTransaction) (txs []blockatlas.Tx) {
	for _, output := range t.Receipt.Outputs {
		vthoTransfers := normalizeVthoTransfers(t, output.Events)
		txs = append(txs, vthoTransfers...)
		transfers := normalizeTxTransfers(t, output.Transfers)
		txs = append(txs, transfers...)
	}
	return txs
}

// normalizeVthoTransfers normalizes vtho transfer events in given transaction
func normalizeVthoTransfers(t *NativeTransaction, events []Event) (txs []blockatlas.Tx) {
	for _, event := range events {
		if len(event.Topics) == 3 && event.Topics[0] == VeThorTransferEvent {
			feeBase10, err := util.HexToDecimal(t.Receipt.Paid)
			if err != nil {
				continue
			}

			valueBase10, err := util.HexToDecimal(event.Data)
			if err != nil {
				continue
			}
			fromHex := event.Topics[1]
			toHex := event.Topics[2]
			from := util.Checksum(formatHexToAddress(fromHex))
			to := util.Checksum(formatHexToAddress(toHex))

			txs = append(txs, blockatlas.Tx{
				ID:       t.ID,
				Coin:     coin.VET,
				From:     from,
				To:       to,
				Fee:      blockatlas.Amount(feeBase10),
				Date:     t.Timestamp,
				Type:     blockatlas.TxNativeTokenTransfer,
				Block:    t.Block,
				Status:   ReceiptStatus(t.Receipt.Reverted),
				Sequence: t.Block,
				Meta: blockatlas.NativeTokenTransfer{
					Name:     GasName,
					Symbol:   GasSymbol,
					TokenID:  GasContract,
					Decimals: 18,
					Value:    blockatlas.Amount(valueBase10),
					From:     from,
					To:       to,
				},
			})
		}
	}
	return
}

// normalizeTxTransfers normalizes transfers in given transaction
func normalizeTxTransfers(t *NativeTransaction, transfers []Transfer) (txs []blockatlas.Tx) {
	for _, transfer := range transfers {
		feeBase10, err := util.HexToDecimal(t.Receipt.Paid)
		if err != nil {
			continue
		}

		valueBase10, err := util.HexToDecimal(transfer.Amount)
		if err != nil {
			continue
		}
		txs = append(txs, blockatlas.Tx{
			ID:       t.ID,
			Coin:     coin.VET,
			From:     util.Checksum(transfer.Sender),
			To:       util.Checksum(transfer.Recipient),
			Fee:      blockatlas.Amount(feeBase10),
			Date:     t.Timestamp,
			Type:     blockatlas.TxTransfer,
			Block:    t.Block,
			Status:   ReceiptStatus(t.Receipt.Reverted),
			Sequence: t.Block,
			Meta: blockatlas.Transfer{
				Value:    blockatlas.Amount(valueBase10),
				Symbol:   coin.Coins[coin.VET].Symbol,
				Decimals: 18,
			},
		})
	}
	return
}

func formatHexToAddress(hex string) string {
	if len(hex) > 26 {
		return "0x" + hex[26:]
	}
	return hex
}

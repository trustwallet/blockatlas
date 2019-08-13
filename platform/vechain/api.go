package vechain

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strings"
	"sync"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("vechain.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.VET]
}

const VeThorContract = "0x0000000000000000000000000000456e65726779"
const VeThorTransferEvent = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

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

	var txs []blockatlas.Tx
	for t := range transactionsChan {
		normalizeTransactionByType(t, func (tx blockatlas.Tx){
			txs = append(txs, tx)
		})
	}

	return &blockatlas.Block{
		Number: num,
		ID:     block.ID,
		Txs:    txs,
	}, nil
}

func normalizeTransactionByType(t *NativeTransaction, append func (tx blockatlas.Tx)) {
	for outputIndex, output := range t.Receipt.Outputs{
		for eventIndex, event := range output.Events {
			if len(event.Topics) == 3 && event.Topics[0] == VeThorTransferEvent {
				if tx, ok := NormalizeTokenTransaction(t, outputIndex, eventIndex); ok {
					append(tx)
				}
			}
		}
		for transferIndex, _ := range output.Transfers {
			if tx, ok := NormalizeTransaction(t, outputIndex, transferIndex); ok {
				append(tx)
			}
		}
	}
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.getTxsByAddress(address)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	if strings.ToLower(token) == VeThorContract {
		return p.getThorTxsByAddress(address)
	} else {
		return nil, nil
	}
}

func (p *Platform) getThorTxsByAddress(address string) ([]blockatlas.Tx, error) {
	sourceTxs, _ := p.client.GetTokenTransfers(address)

	var ids []string
	for _, tx := range sourceTxs.TokenTransfers {
		ids = append(ids, tx.TxID)
	}
	receiptsChan := p.getTransactionReceipt(ids)

	var txs []blockatlas.Tx
	for _, t := range sourceTxs.TokenTransfers {
		if !strings.EqualFold(t.ContractAddress, VeThorContract) {
			continue
		}

		receipt := findTransferReceiptByTxID(receiptsChan, t.TxID)
		if tx, ok := NormalizeTokenTransfer(&t, &receipt); ok {
			txs = append(txs, tx)
		}
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
				logrus.WithError(err).WithField("platform", "vechain").
					Warnf("Failed to get tx receipt for %s", id)
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
				logrus.WithError(err).WithField("platform", "vechain").
					Warnf("Failed to get transaction for %s", id)
			}
			receiptsChan <- receipt
		}(id)
	}

	wg.Wait()
	close(receiptsChan)

	return receiptsChan
}

func findTransferReceiptByTxID(receiptsChan chan *TransferReceipt, txID string) TransferReceipt {

	var transferReceipt TransferReceipt

	for receipt := range receiptsChan {
		if receipt.ID == txID {
			transferReceipt = *receipt
			break
		}
	}

	return transferReceipt
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
			if tx, ok := NormalizeTransfer(receipt, &clause); ok {
				txs = append(txs, tx)
			}
		}
	}

	return txs, nil
}

func formatHexToAddress(hex string) string {
	if len(hex) > 26 {
		return "0x" + hex[26:]
	}
	return hex
}

func NormalizeTransfer(receipt *TransferReceipt, clause *Clause) (tx blockatlas.Tx, ok bool) {
	feeBase10, err := util.HexToDecimal(receipt.Receipt.Paid)
	if err != nil {
		return tx, false
	}
	valueBase10, err := util.HexToDecimal(clause.Value)
	if err != nil {
		return tx, false
	}

	fee := blockatlas.Amount(feeBase10)
	time := receipt.Timestamp
	block := receipt.Block

	return blockatlas.Tx{
		ID:       receipt.ID,
		Coin:     coin.VET,
		From:     receipt.Origin,
		To:       clause.To,
		Fee:      fee,
		Date:     int64(time),
		Type:     blockatlas.TxTransfer,
		Block:    block,
		Status:   ReceiptStatus(receipt.Receipt.Reverted),
		Sequence: block,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(valueBase10),
			Symbol:   coin.Coins[coin.VET].Symbol,
			Decimals: coin.Coins[coin.VET].Decimals,
		},
	}, true
}

func NormalizeTokenTransfer(t *TokenTransfer, receipt *TransferReceipt) (tx blockatlas.Tx, ok bool) {
	feeBase10, err := util.HexToDecimal(receipt.Receipt.Paid)
	if err != nil {
		return tx, false
	}
	valueBase10, err := util.HexToDecimal(t.Amount)
	if err != nil {
		return tx, false
	}
	fee := blockatlas.Amount(feeBase10)
	value := blockatlas.Amount(valueBase10)
	from := t.Origin
	to := t.Receiver
	block := t.Block

	return blockatlas.Tx{
		ID:       t.TxID,
		Coin:     coin.VET,
		From:     from,
		To:       to,
		Fee:      fee,
		Date:     t.Timestamp,
		Type:     blockatlas.TxNativeTokenTransfer,
		Block:    block,
		Status:   ReceiptStatus(receipt.Receipt.Reverted),
		Sequence: block,
		Meta: blockatlas.NativeTokenTransfer{
			Name:     "VeThor Token",
			Symbol:   "VTHO",
			TokenID:  VeThorContract,
			Decimals: 18,
			Value:    value,
			From:     from,
			To:       to,
		},
	}, true
}
// NormalizeTokenTransaction converts a VeChain VTHO token transaction into the generic model
func NormalizeTokenTransaction(t *NativeTransaction, outputIndex int, eventIndex int) (tx blockatlas.Tx, ok bool) {
	feeBase10, err := util.HexToDecimal(t.Receipt.Paid)
	if err != nil {
		return tx, false
	}

	valueBase10, err := util.HexToDecimal(t.Receipt.Outputs[outputIndex].Events[eventIndex].Data)
	if err != nil {
		return tx, false
	}
	fee := blockatlas.Amount(feeBase10)
	value := blockatlas.Amount(valueBase10)
	fromHex := t.Receipt.Outputs[outputIndex].Events[eventIndex].Topics[1]
	toHex := t.Receipt.Outputs[outputIndex].Events[eventIndex].Topics[2]
	from := formatHexToAddress(fromHex)
	to := formatHexToAddress(toHex)
	block := t.Block

	return blockatlas.Tx{
		ID:       t.ID,
		Coin:     coin.VET,
		From:     from,
		To:       to,
		Fee:      fee,
		Date:     t.Timestamp,
		Type:     blockatlas.TxNativeTokenTransfer,
		Block:    block,
		Status:   ReceiptStatus(t.Receipt.Reverted),
		Sequence: block,
		Meta: blockatlas.NativeTokenTransfer{
			Name:     "VeThor Token",
			Symbol:   "VTHO",
			TokenID:  VeThorContract,
			Decimals: 18,
			Value:    value,
			From:     from,
			To:       to,
		},
	}, true
}

// NormalizeTransaction converts a VeChain transaction into the generic model
func NormalizeTransaction(t *NativeTransaction, outputIndex int, transferIndex int) (tx blockatlas.Tx, ok bool) {
	feeBase10, err := util.HexToDecimal(t.Receipt.Paid)
	if err != nil {
		return tx, false
	}

	transfer := t.Receipt.Outputs[outputIndex].Transfers[transferIndex]
	valueBase10, err := util.HexToDecimal(transfer.Amount)
	if err != nil {
		return tx, false
	}

	fee := blockatlas.Amount(feeBase10)
	time := t.Timestamp
	block := t.Block

	return blockatlas.Tx{
		ID:       t.ID,
		Coin:     coin.VET,
		From:     transfer.Sender,
		To:       transfer.Recipient,
		Fee:      fee,
		Date:     time,
		Type:     blockatlas.TxTransfer,
		Block:    block,
		Status:   ReceiptStatus(t.Receipt.Reverted),
		Sequence: block,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(valueBase10),
			Symbol:   coin.Coins[coin.VET].Symbol,
			Decimals: coin.Coins[coin.VET].Decimals,
		},
	}, true
}

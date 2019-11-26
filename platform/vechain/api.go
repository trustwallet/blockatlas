package vechain

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
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

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	cTxs := p.getTransactionsByIDs(block.Transactions)
	txs := make(blockatlas.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return &blockatlas.Block{
		Number: num,
		ID:     block.Id,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
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
	txs := make(blockatlas.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return txs, nil
}

func (p *Platform) getTransactionsByIDs(ids []string) chan blockatlas.TxPage {
	txChan := make(chan blockatlas.TxPage, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(i string, c chan blockatlas.TxPage) {
			defer wg.Done()
			err := p.getTransactionChannel(i, c)
			if err != nil {
				logger.Error(err)
			}
		}(id, txChan)
	}
	wg.Wait()
	close(txChan)
	return txChan
}

func (p *Platform) getTransactionChannel(id string, txChan chan blockatlas.TxPage) error {
	srcTx, err := p.client.GetTransactionByID(id)
	if err != nil {
		return errors.E(err, "Failed to get tx", errors.TypePlatformUnmarshal,
			errors.Params{"id": id}).PushToSentry()
	}

	receipt, err := p.client.GetTransactionReceiptByID(id)
	if err != nil {
		return errors.E(err, "Failed to get tx id receipt", errors.TypePlatformUnmarshal,
			errors.Params{"id": id}).PushToSentry()
	}

	txs, err := NormalizeTokenTransaction(srcTx, receipt)
	if err != nil {
		return errors.E(err, "Failed to NormalizeBlockTransactions tx", errors.TypePlatformUnmarshal,
			errors.Params{"tx": srcTx}).PushToSentry()
	}
	txChan <- txs
	return nil
}

func NormalizeTokenTransaction(srcTx Tx, receipt TxReceipt) (blockatlas.TxPage, error) {
	if receipt.Outputs == nil || len(receipt.Outputs) == 0 {
		return blockatlas.TxPage{}, errors.E("NormalizeBlockTransaction: Clauses not found", errors.Params{"tx": srcTx}).PushToSentry()
	}

	nonce, err := hexToInt(srcTx.Nonce)
	if err != nil {
		return blockatlas.TxPage{}, err
	}
	origin := util.GetValidParameter(srcTx.Origin, srcTx.Meta.TxOrigin)

	fee, err := util.HexToDecimal(receipt.Paid)
	if err != nil {
		return blockatlas.TxPage{}, err
	}

	txs := make(blockatlas.TxPage, 0)
	for _, output := range receipt.Outputs {
		event := output.Events[0] // TODO add support for multisend
		value, err := util.HexToDecimal(event.Data)
		if err != nil {
			return blockatlas.TxPage{}, err
		}

		txs = append(txs, blockatlas.Tx{
			ID:       srcTx.Id,
			Coin:     coin.VET,
			From:     origin,
			To:       event.Address,
			Fee:      blockatlas.Amount(fee),
			Date:     srcTx.Meta.BlockTimestamp,
			Type:     blockatlas.TxTokenTransfer,
			Block:    srcTx.Meta.BlockNumber,
			Sequence: nonce,
			Status:   blockatlas.StatusCompleted,
			Meta: blockatlas.TokenTransfer{
				Name: "",
				TokenID: "",
				Value:    blockatlas.Amount(value),
				Symbol:   "VTHO", // TODO replace with real name for other coins
				Decimals: 18, // TODO Not all tokens have decimal 18 https://github.com/vechain/token-registry/tree/master/tokens/main
				From: origin,
				To: getRecipientAddress(event.Topics[2]),
			},
		})
	}
	return txs, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	headBlock, err := p.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}
	transfers, err := p.client.GetTransactions(address, headBlock)
	if err != nil {
		return nil, err
	}

	txs := make(blockatlas.TxPage, 0)
	for _, t := range transfers {
		trxId, err := p.client.GetTransactionByID(t.Meta.TxId)
		tx, err := NormalizeTransaction(t, trxId)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func NormalizeTransaction(srcTx LogTransfer, trxId Tx) (blockatlas.Tx, error) {
	value, err := util.HexToDecimal(srcTx.Amount)
	if err != nil {
		return blockatlas.Tx{}, err
	}

	nonce, err := hexToInt(trxId.Nonce)
	if err != nil {
		return blockatlas.Tx{}, err
	}
	fee := strconv.Itoa(trxId.Gas)

	return blockatlas.Tx{
		ID:     srcTx.Meta.TxId,
		Coin:   coin.VET,
		From:   srcTx.Sender,
		To:     srcTx.Recipient,
		Fee:    blockatlas.Amount(fee),
		Date:   srcTx.Meta.BlockTimestamp,
		Type:   blockatlas.TxTransfer,
		Block:  srcTx.Meta.BlockNumber,
		Status: blockatlas.StatusCompleted,
		Sequence: nonce,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[coin.VET].Symbol,
			Decimals: 18,
		},
	}, nil
}

func hexToInt(hex string) (uint64, error) {
	nonceStr, err := util.HexToDecimal(hex)
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

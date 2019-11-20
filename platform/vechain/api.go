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
	cTxs := p.getTransactions(block.Transactions)
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

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	num, err := p.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}
	tks, err := p.client.GetTokens(address, token, num)
	if err != nil {
		return nil, err
	}
	logTxs := make([]string, 0)
	for _, t := range tks {
		logTxs = append(logTxs, t.Meta.TxId)
	}

	cTxs := p.getTransactions(logTxs)
	txs := make(blockatlas.TxPage, 0)
	for t := range cTxs {
		txs = append(txs, t...)
	}
	return txs, nil
}

func (p *Platform) getTransactions(ids []string) chan blockatlas.TxPage {
	txChan := make(chan blockatlas.TxPage, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := p.getTransactionChannel(id, txChan)
			if err != nil {
				logger.Error(err)
			}
		}()
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
	txs, err := NormalizeTransaction(srcTx)
	if err != nil {
		return errors.E(err, "Failed to NormalizeBlockTransactions tx", errors.TypePlatformUnmarshal,
			errors.Params{"tx": srcTx}).PushToSentry()
	}
	txChan <- txs
	return nil
}

func NormalizeTransaction(srcTx Tx) (blockatlas.TxPage, error) {
	if srcTx.Clauses == nil || len(srcTx.Clauses) == 0 {
		return blockatlas.TxPage{}, errors.E("NormalizeBlockTransaction: Clauses not found", errors.Params{"tx": srcTx}).PushToSentry()
	}

	nonce, err := hexToInt(srcTx.Nonce)
	if err != nil {
		return blockatlas.TxPage{}, err
	}

	id := util.GetValidParameter(srcTx.Id, srcTx.Meta.TxId)
	origin := util.GetValidParameter(srcTx.Origin, srcTx.Meta.TxOrigin)
	fee := strconv.Itoa(srcTx.Gas)

	txs := make(blockatlas.TxPage, 0)
	for _, clause := range srcTx.Clauses {
		value, err := util.HexToDecimal(clause.Value)
		if err != nil {
			return blockatlas.TxPage{}, err
		}

		txs = append(txs, blockatlas.Tx{
			ID:       id,
			Coin:     coin.VET,
			From:     origin,
			To:       clause.To,
			Fee:      blockatlas.Amount(fee),
			Date:     srcTx.Meta.BlockTimestamp,
			Type:     blockatlas.TxTransfer,
			Block:    srcTx.Meta.BlockNumber,
			Sequence: uint64(nonce),
			Status:   blockatlas.StatusCompleted,
			Meta: blockatlas.Transfer{
				Value:    blockatlas.Amount(value),
				Symbol:   coin.Coins[coin.VET].Symbol,
				Decimals: 18,
			},
		})
	}
	return txs, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	num, err := p.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}
	srcTxs, err := p.client.GetTransactions(address, num)
	if err != nil {
		return nil, err
	}

	txs := make(blockatlas.TxPage, 0)
	for _, t := range srcTxs {
		tx, err := NormalizeLogTransaction(t)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func NormalizeLogTransaction(srcTx LogTx) (blockatlas.Tx, error) {
	value, err := util.HexToDecimal(srcTx.Amount)
	if err != nil {
		return blockatlas.Tx{}, err
	}

	id := util.GetValidParameter(srcTx.Id, srcTx.Meta.TxId)
	tx := blockatlas.Tx{
		ID:     id,
		Coin:   coin.VET,
		From:   srcTx.Sender,
		To:     srcTx.Recipient,
		Fee:    blockatlas.Amount("0"),
		Date:   srcTx.Meta.BlockTimestamp,
		Type:   blockatlas.TxTransfer,
		Block:  srcTx.Meta.BlockNumber,
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[coin.VET].Symbol,
			Decimals: 18,
		},
	}
	return tx, nil
}

func hexToInt(hex string) (int64, error) {
	nonceStr, err := util.HexToDecimal(hex)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(nonceStr, 10, 64)
}

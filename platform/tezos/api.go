package tezos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

const Annual = 6.09

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("tezos.api"))}
	p.client.SetTimeout(35)
	p.rpcClient = RpcClient{blockatlas.InitClient(viper.GetString("tezos.rpc"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txTypes := []TxType{TxTransactions}
	var wg sync.WaitGroup
	out := make(chan []Transaction, len(txTypes))
	wg.Add(len(txTypes))
	for _, t := range txTypes {
		go func(txType TxType, addr string, wg *sync.WaitGroup) {
			defer wg.Done()
			txs, err := p.client.GetTxsOfAddress(address, txType)
			if err != nil {
				logger.Error("GetAddrTxs", err, logger.Params{"txType": txType, "addr": addr})
				return
			}
			out <- txs
		}(t, address, &wg)
	}
	wg.Wait()
	close(out)
	srcTxs := make([]Transaction, 0)
	for r := range out {
		srcTxs = append(srcTxs, r...)
	}
	return NormalizeTxs(srcTxs), nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	txTypes := []TxType{TxTransactions, TxDelegations}
	var wg sync.WaitGroup
	out := make(chan []Transaction, len(txTypes))
	wg.Add(len(txTypes))
	for _, t := range txTypes {
		go func(txType TxType, num int64, wg *sync.WaitGroup) {
			defer wg.Done()
			txs, err := p.client.GetBlockByNumber(num, txType)
			if err != nil {
				logger.Error("GetAddrTxs", err, logger.Params{"txType": txType, "num": num})
				return
			}
			out <- txs
		}(t, num, &wg)
	}
	wg.Wait()
	close(out)
	srcTxs := make([]Transaction, 0)
	for r := range out {
		srcTxs = append(srcTxs, r...)
	}
	txs := NormalizeTxs(srcTxs)
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func NormalizeTxs(srcTxs []Transaction) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

// NormalizeTx converts a Tezos transaction into the generic model
func NormalizeTx(srcTx Transaction) (blockatlas.Tx, bool) {
	errMsg := ""
	status := blockatlas.StatusCompleted
	if srcTx.Status() != TxStatusApplied {
		errMsg = "transaction failed"
		status = blockatlas.StatusFailed
	}
	tx := blockatlas.Tx{
		ID:     srcTx.Op.OpHash,
		Coin:   coin.XTZ,
		Date:   srcTx.Op.BlockTimestamp.Unix(),
		From:   srcTx.Source(),
		To:     srcTx.Destination(),
		Fee:    blockatlas.Amount(srcTx.Fee()),
		Block:  srcTx.Op.BlockLevel,
		Status: status,
		Error:  errMsg,
	}

	switch srcTx.Kind() {
	case TxKindDelegation:
		title := blockatlas.AnyActionDelegation
		if len(srcTx.Delegation.Delegate) == 0 {
			title = blockatlas.AnyActionUndelegation
		}
		tx.Meta = blockatlas.AnyAction{
			Coin:     coin.Tezos().ID,
			Title:    title,
			Key:      blockatlas.KeyStakeDelegate,
			Name:     coin.Tezos().Name,
			Symbol:   coin.Tezos().Symbol,
			Decimals: coin.Tezos().Decimals,
		}
	case TxKindTransaction:
		tx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Tx.Amount),
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		}
	default:
		return blockatlas.Tx{}, false
	}
	return tx, true
}

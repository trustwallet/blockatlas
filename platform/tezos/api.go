package tezos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	stakeClient    Client
	rpcClient RpcClient
}

const Annual = 6.09

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("tezos.api"))}
	p.client.SetTimeout(30)
	p.stakeClient = Client{blockatlas.InitClient(viper.GetString("tezos.stake_api"))}
	p.rpcClient = RpcClient{blockatlas.InitClient(viper.GetString("tezos.rpc"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(s)
	return txs, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcBlock)
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
func NormalizeTx(srcTx Transaction) (tx blockatlas.Tx, ok bool) {
	if srcTx.Tx.Kind != "transaction" {
		return tx, false
	}

	var status blockatlas.Status
	var errMsg string
	if srcTx.Tx.Status == "applied" {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
		errMsg = "transaction failed"
	}
	return blockatlas.Tx{
		ID:    srcTx.Op.OpHash,
		Coin:  coin.XTZ,
		Date:  srcTx.Op.BlockTimestamp.Unix(),
		From:  srcTx.Tx.Source,
		To:    srcTx.Tx.Destination,
		Fee:   blockatlas.Amount(srcTx.Tx.Fee),
		Block: srcTx.Op.BlockLevel,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Tx.Amount),
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		},
		Status: status,
		Error:  errMsg,
	}, true
}

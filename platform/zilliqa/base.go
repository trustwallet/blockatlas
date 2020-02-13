package zilliqa

import (
	"strconv"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"

	"github.com/spf13/viper"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
	udClient  Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("zilliqa.api"))}
	p.client.Headers["X-APIKEY"] = viper.GetString("zilliqa.key")

	p.rpcClient = RpcClient{blockatlas.InitClient(viper.GetString("zilliqa.rpc"))}
	p.udClient = Client{blockatlas.InitClient(viper.GetString("zilliqa.lookup"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ZIL]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	info, err := p.rpcClient.GetBlockchainInfo()
	if err != nil {
		return 0, err
	}
	block, err := strconv.ParseInt(info.NumTxBlocks, 10, 64)
	if err != nil {
		return 0, err
	}

	return block, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	var normalized []blockatlas.Tx
	txs, err := p.rpcClient.GetTxInBlock(num)
	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		normalized = append(normalized, tx)
	}
	block := blockatlas.Block{
		Number: num,
		Txs:    normalized,
	}

	return &block, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	var normalized []blockatlas.Tx
	txs, err := p.client.GetTxsOfAddress(address)

	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		if len(normalized) >= blockatlas.TxPerPage {
			break
		}
		normalized = append(normalized, tx)
	}

	return normalized, nil
}

func Normalize(srcTx *Tx) (tx blockatlas.Tx) {
	tx = blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.ZIL,
		Date:     srcTx.Timestamp / 1000,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      blockatlas.Amount(srcTx.Fee),
		Block:    srcTx.BlockHeight,
		Sequence: srcTx.NonceValue(),
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.ZIL].Symbol,
			Decimals: coin.Coins[coin.ZIL].Decimals,
		},
	}
	if !srcTx.ReceiptSuccess {
		tx.Status = blockatlas.StatusFailed
	}
	return tx
}

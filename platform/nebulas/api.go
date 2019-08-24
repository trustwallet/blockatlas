package nebulas

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"math/big"
	"net/http"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("nebulas.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.NAS]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txs, err := p.client.GetTxs(address, 1)
	if err != nil {
		return nil, err
	}

	var normalizeTxs []blockatlas.Tx
	for _, srcTx := range txs {
		normalizeTxs = append(normalizeTxs, NormalizeTx(srcTx))
	}
	return normalizeTxs, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestIrreversibleBlock()
}

// The below method would probably be the `GetTxsByAddress`.
func (p *Platform) GetTxsByHash(hash string) (blockatlas.TxPage, error) {
	nebulaBlock, err := p.client.GetBlockByHash(hash, true)

	return handleBlockResponse(nebulaBlock, err)
}

func (p *Platform) GetBlockByNumber(height int64) (blockatlas.TxPage, error) {
	nebulaBlock, err := p.client.GetBlockByHeight(height, true)

	return handleBlockResponse(nebulaBlock, err)
}

func handleBlockResponse(block NebulaBlock, err error) (blockatlas.TxPage, error) {
	if err != nil {
		return nil, err
	}
	var normalizeTxs []blockatlas.Tx
	for _, srcTx := range block.Transactions {
		normalizeTxs = append(normalizeTxs, NormalizeNebulaTx(srcTx, block.Height))
	}
	return normalizeTxs, nil
}

func NormalizeTx(srcTx Transaction) blockatlas.Tx {
	var status = blockatlas.StatusCompleted
	if srcTx.Status == 0 {
		status = blockatlas.StatusFailed
	}
	return blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.NAS,
		From:     srcTx.From.Hash,
		To:       srcTx.To.Hash,
		Fee:      blockatlas.Amount(srcTx.TxFee),
		Date:     int64(srcTx.Timestamp) / 1000,
		Block:    srcTx.Block.Height,
		Status:   status,
		Sequence: srcTx.Nonce,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.NAS].Symbol,
			Decimals: coin.Coins[coin.NAS].Decimals,
		},
	}
}

func calculateTxFee(gasPrice string, gasUsed string) string {
	var gPrice, gUsed, fee big.Int
	// Using big int because they are string values and also can be really large,
	// so big Int is the fastest approach
	gPrice.SetString(gasPrice, 10)
	gUsed.SetString(gasUsed, 10)

	fee.Mul(&gPrice, &gUsed)
	return fee.String()
}

func parseToUint64(source string) uint64 {
	result, err := strconv.ParseUint(source, 10, 64)
	if err != nil{
		return 0
	}
	return result
}

func parseToInt(source string) int64 {
	result, err := strconv.ParseInt(source, 10, 64)
	if err != nil{
		return  0
	}
	return result
}

func NormalizeNebulaTx(srcTx NebulaTransaction, nebulaBlockHeight string) blockatlas.Tx {
	var status = blockatlas.StatusCompleted
	if srcTx.Status == 0 {
		status = blockatlas.StatusFailed
	}

	time := parseToInt(srcTx.Timestamp)

	block := parseToUint64(nebulaBlockHeight)
	nonce := parseToUint64(srcTx.Nonce)

	txFee := calculateTxFee(srcTx.GasPrice, srcTx.GasUsed)

	return blockatlas.Tx{
		ID:        srcTx.Hash,
		Coin:      coin.NAS,
		From:      srcTx.From,
		To:        srcTx.To,
		Fee:       blockatlas.Amount(txFee),
		Date:      time,
		Block:     block,
		Status:    status,
		Sequence:  nonce,
		Meta:      blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.NAS].Symbol,
			Decimals: coin.Coins[coin.NAS].Decimals,
		},
	}
}

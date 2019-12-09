package harmony

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("harmony.rpc"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ONE]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	result, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return []blockatlas.Tx{}, err
	}
	return NormalizeTxs(result.Transactions), err
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func NormalizeTxs(txs []Transaction) []blockatlas.Tx {
	normalizeTxs := make([]blockatlas.Tx, 0)
	for _, srcTx := range txs {
		normalized, isCorrect := NormalizeTx(&srcTx)
		if !isCorrect {
			return []blockatlas.Tx{}
		}
		normalizeTxs = append(normalizeTxs, normalized)
	}
	return normalizeTxs
}

func NormalizeTx(trx *Transaction) (tx blockatlas.Tx, b bool) {
	gasPrice, err := hexToInt(trx.GasPrice)
	gas, err := hexToInt(trx.Gas)
	fee := gas * gasPrice
	literalFee := strconv.Itoa(int(fee))

	value, err := hexToInt(trx.Value)
	literalValue := strconv.Itoa(int(value))

	block, err := hexToInt(trx.BlockNumber)

	if err != nil {
		return blockatlas.Tx{}, false
	}

	return blockatlas.Tx{
		ID:     trx.Hash,
		Coin:   coin.ONE,
		From:   trx.From,
		To:     trx.To,
		Fee:    blockatlas.Amount(literalFee),
		Status: blockatlas.StatusCompleted,
		Date:   0,
		Type:   blockatlas.TxTransfer,
		Block:  block,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(literalValue),
			Symbol:   coin.Coins[coin.ONE].Symbol,
			Decimals: coin.Coins[coin.ONE].Decimals,
		},
	}, true
}

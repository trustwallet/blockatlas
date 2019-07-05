package semux

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("semux.api")
	p.client.Init()
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.SEM]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range s {
		tx, err := Normalize(&srcTx)
		if err == nil {
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

// Normalize converts a semux transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, err error) {
	blockNumber, err := strconv.ParseUint(srcTx.BlockNumber, 10, 64)
	if err != nil {
		logrus.Error("Failed to convert TX blockNumber for Semux API")
		return blockatlas.Tx{}, err
	}

	date, err := strconv.ParseInt(srcTx.Timestamp, 10, 64)
	if err != nil {
		logrus.Error("Failed to convert TX timestamp for Semux API")
		return blockatlas.Tx{}, err
	}

	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.SEM,
		Date:  date / 1000,
		From:  srcTx.From,
		To:    srcTx.To,
		Fee:   srcTx.Fee,
		Block: blockNumber,
		Meta: blockatlas.Transfer{
			Value: srcTx.Value,
		},
	}, nil
}

package polkadot

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString(p.ConfigKey()))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) ConfigKey() string {
	return fmt.Sprintf("%s.api", p.Coin().Handle)
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTransfersOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range s {
		tx := p.NormalizeTransfer(&srcTx)
		txs = append(txs, tx)
	}

	return txs, nil
}

func (p *Platform) NormalizeTransfer(srcTx *Transfer) blockatlas.Tx {
	result := blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  p.Coin().ID,
		Date:  int64(srcTx.BlockTimestamp),
		From:  srcTx.From,
		To:    srcTx.To,
		Fee:   "100000000", // API will return fee later
		Block: srcTx.BlockNum,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Amount),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}
	return result
}

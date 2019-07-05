package waves

import (
	"net/http"
	"strconv"

	"github.com/trustwallet/blockatlas/coin"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("waves.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.WAVES]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, 25)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range addressTxs {
		txs = append(txs, NormalizeTx(&srcTx, p.Coin().ID))
	}

	return txs, nil
}

func NormalizeTx(srcTx *Transaction, coinIndex uint) blockatlas.Tx {
	return blockatlas.Tx{
		ID:     srcTx.Id,
		Coin:   coinIndex,
		From:   srcTx.Sender,
		To:     srcTx.Recipient,
		Fee:    blockatlas.Amount(strconv.Itoa(int(srcTx.Fee))),
		Date:   int64(srcTx.Timestamp) / 1000,
		Block:  srcTx.Block,
		Memo:   srcTx.Attachment,
		Status: blockatlas.StatusCompleted,
		Meta:   blockatlas.Transfer{
			Value: blockatlas.Amount(strconv.Itoa(int(srcTx.Amount))),
		},
	}
}

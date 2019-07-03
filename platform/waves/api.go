package waves

import (
	"net/http"
	"strconv"

	"github.com/trustwallet/blockatlas/coin"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
)

const Handle = "waves"

type Platform struct {
	client Client
}

func (p *Platform) Handle() string {
	return Handle
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
		tx, ok := NormalizeTx(&srcTx, p.Coin().ID)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func NormalizeTx(srcTx *Transaction, coinIndex uint) (tx blockatlas.Tx, ok bool) {
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return tx, false
	}
	baseTx.Meta = blockatlas.Transfer{
		Value: blockatlas.Amount(strconv.Itoa(int(srcTx.Amount))),
	}
	return tx, true
}

func extractBase(srcTx *Transaction, coinIndex uint) (base blockatlas.Tx, ok bool) {
	base = blockatlas.Tx{
		ID:     srcTx.Id,
		Coin:   coinIndex,
		From:   srcTx.Sender,
		To:     srcTx.Recipient,
		Fee:    blockatlas.Amount(strconv.Itoa(int(srcTx.Fee))),
		Date:   int64(srcTx.Timestamp) / 1000,
		Block:  srcTx.Block,
		Memo:   srcTx.Attachment,
		Status: blockatlas.StatusCompleted,
	}
	return base, true
}

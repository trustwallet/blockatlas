package icon

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"time"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.RPCURL = viper.GetString("icon.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ICX]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	trxs, err := p.client.GetAddressTransactions(address)
	if err != nil {
		return nil, err
	}

	nTrxs := make([]blockatlas.Tx, 0)
	for _, trx := range trxs {
		nTrx, ok := Normalize(&trx)
		if !ok {
			continue
		}
		nTrxs = append(nTrxs, nTrx)
	}

	return nTrxs, nil
}

// Normalize converts an Icon transaction into the generic model
func Normalize(trx *Tx) (tx blockatlas.Tx, b bool) {
	date, err := time.Parse("2006-01-02T15:04:05.999Z0700", trx.CreateDate)
	if err != nil {
		fmt.Printf("%v\n", err)
		return tx, false
	}
	fee := util.DecimalExp(string(trx.Fee), 18)
	value := util.DecimalExp(string(trx.Amount), 18)

	return blockatlas.Tx{
		ID:      trx.TxHash,
		Coin   : coin.ICX,
		From   : trx.FromAddr,
		To     : trx.ToAddr,
		Fee    : blockatlas.Amount(fee),
		Status : blockatlas.StatusCompleted,
		Date   : date.Unix(),
		Type   : blockatlas.TxTransfer,
		Block  : trx.Height,
		Meta: blockatlas.Transfer{
			Value : blockatlas.Amount(value),
		},
	}, true
}

package ripple

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/valyala/fastjson"
	"net/http"
	"time"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("ripple.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XRP]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range s {
		tx, ok := Normalize(&srcTx)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

// Normalize converts a Ripple transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	// Only accept XRP payments (typeof tx.amount === 'string')
	var p fastjson.Parser
	v, pErr := p.ParseBytes(srcTx.Payment.Amount)
	if pErr != nil {
		return tx, false
	}
	if v.Type() != fastjson.TypeString {
		return tx, false
	}
	srcAmount := string(v.GetStringBytes())

	date, err := time.Parse("2006-01-02T15:04:05-07:00", srcTx.Date)
	var unix int64
	if err != nil {
		unix = 0
	} else {
		unix = date.Unix()
	}

	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.XRP,
		Date:  unix,
		From:  srcTx.Payment.Account,
		To:    srcTx.Payment.Destination,
		Fee:   srcTx.Payment.Fee,
		Block: srcTx.LedgerIndex,
		Meta:  blockatlas.Transfer{
			Value: blockatlas.Amount(srcAmount),
		},
	}, true
}

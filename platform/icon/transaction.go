package icon

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	trxs, err := p.client.GetAddressTransactions(address)
	if err != nil {
		return nil, err
	}

	nTrxs := make([]txtype.Tx, 0)
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
func Normalize(trx *Tx) (tx txtype.Tx, b bool) {
	date, err := time.Parse("2006-01-02T15:04:05.999Z0700", trx.CreateDate)
	if err != nil {
		log.Error(err)
		return tx, false
	}
	fee := numbers.DecimalExp(string(trx.Fee), 18)
	value := numbers.DecimalExp(string(trx.Amount), 18)

	return txtype.Tx{
		ID:     trx.TxHash,
		Coin:   coin.ICX,
		From:   trx.FromAddr,
		To:     trx.ToAddr,
		Fee:    txtype.Amount(fee),
		Status: txtype.StatusCompleted,
		Date:   date.Unix(),
		Type:   txtype.TxTransfer,
		Block:  trx.Height,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(value),
			Symbol:   coin.Coins[coin.ICX].Symbol,
			Decimals: coin.Coins[coin.ICX].Decimals,
		},
	}, true
}

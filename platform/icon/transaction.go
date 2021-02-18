package icon

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	trxs, err := p.client.GetAddressTransactions(address)
	if err != nil {
		return nil, err
	}

	nTrxs := make(types.Txs, 0)
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
func Normalize(trx *Tx) (tx types.Tx, b bool) {
	date, err := time.Parse("2006-01-02T15:04:05.999Z0700", trx.CreateDate)
	if err != nil {
		log.Error(err)
		return tx, false
	}
	fee := numbers.DecimalExp(string(trx.Fee), 18)
	value := numbers.DecimalExp(string(trx.Amount), 18)

	return types.Tx{
		ID:     trx.TxHash,
		Coin:   coin.ICON,
		From:   trx.FromAddr,
		To:     trx.ToAddr,
		Fee:    types.Amount(fee),
		Status: types.StatusCompleted,
		Date:   date.Unix(),
		Type:   types.TxTransfer,
		Block:  trx.Height,
		Meta: types.Transfer{
			Value:    types.Amount(value),
			Symbol:   coin.Icon().Symbol,
			Decimals: coin.Icon().Decimals,
		},
	}, true
}

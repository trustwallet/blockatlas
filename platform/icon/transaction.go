package icon

import (
	"time"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
)

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
		err = errors.E(err, errors.TypePlatformUnmarshal)
		logger.Error(err)
		return tx, false
	}
	fee := numbers.DecimalExp(string(trx.Fee), 18)
	value := numbers.DecimalExp(string(trx.Amount), 18)

	return blockatlas.Tx{
		ID:     trx.TxHash,
		Coin:   coin.ICON,
		From:   trx.FromAddr,
		To:     trx.ToAddr,
		Fee:    blockatlas.Amount(fee),
		Status: blockatlas.StatusCompleted,
		Date:   date.Unix(),
		Type:   blockatlas.TxTransfer,
		Block:  trx.Height,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[coin.ICON].Symbol,
			Decimals: coin.Coins[coin.ICON].Decimals,
		},
	}, true
}

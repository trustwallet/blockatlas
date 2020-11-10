package stellar

import (
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	payments, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	return p.NormalizePayments(payments), nil
}

func (p *Platform) NormalizePayments(payments []Payment) []blockatlas.Tx {
	txs := make([]blockatlas.Tx, 0, len(payments))
	for _, payment := range payments {
		if tx, ok := Normalize(&payment, p.CoinIndex); ok {
			txs = append(txs, tx)
		}
	}
	return txs
}

// Normalize converts a Stellar-based transaction into the generic model
func Normalize(payment *Payment, nativeCoinIndex uint) (tx blockatlas.Tx, ok bool) {
	switch payment.Type {
	case PaymentType:
		if payment.AssetType != Native {
			return tx, false
		}
	case CreateAccount:
		break
	default:
		return tx, false
	}
	date, err := time.Parse("2006-01-02T15:04:05Z", payment.CreatedAt)
	if err != nil {
		return tx, false
	}
	var value, from, to string
	if payment.Amount != "" {
		value, err = numbers.DecimalToSatoshis(payment.Amount)
		from = payment.From
		to = payment.To
	} else if payment.StartingBalance != "" { // When transfer to new account
		value, err = numbers.DecimalToSatoshis(payment.StartingBalance)
		from = payment.Funder
		to = payment.Account
	} else {
		return tx, false
	}
	if err != nil {
		return tx, false
	}
	return blockatlas.Tx{
		ID:    payment.TransactionHash,
		Coin:  nativeCoinIndex,
		From:  from,
		To:    to,
		Fee:   FixedFee,
		Date:  date.Unix(),
		Memo:  payment.Transaction.Memo,
		Block: payment.Transaction.Ledger,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[nativeCoinIndex].Symbol,
			Decimals: coin.Coins[nativeCoinIndex].Decimals,
		},
	}, true
}

package polkadot

import (
	"math"
	"math/big"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func ParseAmount(s string, decimals uint) blockatlas.Amount {
	base := big.NewFloat(float64(math.Pow10(int(decimals))))
	bigFloat, _, err := big.ParseFloat(s, 10, decimals*10, big.ToNearestEven)
	var amount string
	if err == nil {
		bigFloat = bigFloat.Mul(bigFloat, base)
		amount = bigFloat.Text('f', 0)
	} else {
		amount = "0"
	}
	return blockatlas.Amount(amount)
}

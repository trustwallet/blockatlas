package blockbook

import (
	"math/big"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (s *EthereumSpecific) GetStatus() (blockatlas.Status, string) {
	switch s.Status {
	case -1:
		return blockatlas.StatusPending, ""
	case 0:
		return blockatlas.StatusError, "Error"
	case 1:
		return blockatlas.StatusCompleted, ""
	default:
		return blockatlas.StatusError, "Unable to define transaction status"
	}
}

func (t *Transaction) FromAddress() string {
	if len(t.Vin) > 0 && len(t.Vin[0].Addresses) > 0 {
		return t.Vin[0].Addresses[0]
	}
	return ""
}

func (t *Transaction) GetFee() string {
	status, _ := t.EthereumSpecific.GetStatus()
	if status != blockatlas.StatusPending {
		return t.Fees
	}

	gasLimit := t.EthereumSpecific.GasLimit
	gasPrice, ok := new(big.Int).SetString(t.EthereumSpecific.GasPrice, 10)
	if gasLimit == nil || !ok {
		return "0"
	}
	fee := new(big.Int).Mul(gasLimit, gasPrice)
	return fee.String()
}

func (t *Transaction) ToAddress() string {
	if len(t.Vout) > 0 && len(t.Vout[0].Addresses) > 0 {
		return t.Vout[0].Addresses[0]
	}
	return ""
}

func GetDirection(address, from, to string) blockatlas.Direction {
	if address == from && address == to {
		return blockatlas.DirectionSelf
	}
	if address == from {
		return blockatlas.DirectionOutgoing
	}
	return blockatlas.DirectionIncoming
}

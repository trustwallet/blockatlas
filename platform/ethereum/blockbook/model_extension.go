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

func (transaction *Transaction) FromAddress() string {
	if len(transaction.Vin) > 0 && len(transaction.Vin[0].Addresses) > 0 {
		return transaction.Vin[0].Addresses[0]
	}
	return ""
}

func (transaction *Transaction) GetFee() string {
	status, _ := transaction.EthereumSpecific.GetStatus()
	if status != blockatlas.StatusPending {
		return transaction.Fees
	}

	gasLimit := transaction.EthereumSpecific.GasLimit
	gasPrice, ok := new(big.Int).SetString(transaction.EthereumSpecific.GasPrice, 10)
	if gasLimit == nil || !ok {
		return "0"
	}
	fee := new(big.Int).Mul(gasLimit, gasPrice)
	return fee.String()
}

func (transaction *Transaction) ToAddress() string {
	if len(transaction.Vout) > 0 && len(transaction.Vout[0].Addresses) > 0 {
		return transaction.Vout[0].Addresses[0]
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

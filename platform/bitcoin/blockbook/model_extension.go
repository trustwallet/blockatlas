package blockbook

import (
	"math/big"

	"github.com/trustwallet/golibs/types"
)

func (s *EthereumSpecific) GetStatus() (types.Status, string) {
	switch s.Status {
	case -1:
		return types.StatusPending, ""
	case 0:
		return types.StatusError, "Error"
	case 1:
		return types.StatusCompleted, ""
	default:
		return types.StatusError, "Unable to define transaction status"
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
	if status != types.StatusPending {
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

func GetDirection(address, from, to string) types.Direction {
	if address == from && address == to {
		return types.DirectionSelf
	}
	if address == from {
		return types.DirectionOutgoing
	}
	return types.DirectionIncoming
}

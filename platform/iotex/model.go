package iotex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Response struct {
	ActionInfo []*ActionInfo `json:"actionInfo"`
}

type AccountInfo struct {
	AccountMeta *AccountMeta `json:"accountMeta"`
}

type AccountMeta struct {
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	Nonce        string `json:"nonce"`
	PendingNonce string `json:"pendingNonce"`
	NumActions   string `json:"numActions"`
}

type ActionInfo struct {
	Action    *Action `json:"action"`
	ActHash   string  `json:"actHash"`
	BlkHeight string  `json:"blkHeight"`
	Sender    string  `json:"sender"`
	GasFee    string  `json:"gasFee"`
	Timestamp string  `json:"timestamp"`
}

type Action struct {
	Core *ActionCore `json:"core"`
}

type ActionCore struct {
	Nonce    string    `json:"nonce"`
	Transfer *Transfer `json:"transfer"`
}

type Transfer struct {
	Amount    blockatlas.Amount `json:"amount"`
	Recipient string            `json:"recipient"`
}

type ChainMeta struct {
	Height string `json:"height"`
}

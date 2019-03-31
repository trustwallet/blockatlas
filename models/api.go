package models

import (
	"github.com/spf13/cast"
	"sort"
)

const (
	TxTransfer            = "transfer"
	TxTokenTransfer       = "token_transfer"
	TxCollectibleTransfer = "collectible_transfer"
	TxTokenSwap           = "token_swap"
	TxContractCall        = "contract_call"
)

const TxPerPage = 25

type Response []Tx

func (r *Response) Sort() {
	sort.Slice(*r, func(i, j int) bool {
		ti := cast.ToUint64((*r)[i].Date)
		tj := cast.ToUint64((*r)[j].Date)
		return ti >= tj
	})
}

type Amount string

type Tx struct {
	Id    string      `json:"id"`
	Coin  uint        `json:"coin"`
	From  string      `json:"from"`
	To    string      `json:"to"`
	Fee   Amount      `json:"fee"`
	Date  int64       `json:"date"`
	Type  string      `json:"type"`
	Block uint64      `json:"block,omitempty"`
	Meta  interface{} `json:"metadata"`
}

type Transfer struct {
	Value Amount `json:"value"`
}

type TokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Contract string `json:"contract"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
}

type CollectibleTransfer struct {
	Name     string `json:"name"`
	Contract string `json:"contract"`
	ImageUrl string `json:"image_url"`
}

type TokenSwap struct {
	Input  TokenTransfer `json:"input"`
	Output TokenTransfer `json:"output"`
}

type ContractCall struct {
	// TODO
}

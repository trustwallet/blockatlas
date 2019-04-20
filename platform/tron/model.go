package tron

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/models"
)

type Page struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
	Txs []Tx `json:"data"`
}

type Tx struct {
	ID string `json:"txID"`
	Data TxData `json:"raw_data"`
}

type TxData struct {
	Contracts []Contract `json:"contract"`
	Timestamp int64 `json:"timestamp"`
}

type Contract struct {
	Type string `json:"type"`
	Parameter interface{} `json:"parameter"`
}

type TransferContract struct {
	Value TransferValue `json:"value"`
}

type TransferValue struct {
	Amount models.Amount `json:"amount"`
	OwnerAddress string `json:"owner_address"`
	ToAddress string `json:"to_address"`
}

func (c *Contract) UnmarshalJSON(buf []byte) error {
	var contractInternal struct {
		Type string `json:"type"`
		Parameter json.RawMessage `json:"parameter"`
	}
	err := json.Unmarshal(buf, &contractInternal)
	if err != nil {
		return err
	}
	switch contractInternal.Type {
	case "TransferContract":
		var transfer TransferContract
		err = json.Unmarshal(contractInternal.Parameter, &transfer)
		c.Parameter = transfer
	}
	return err
}

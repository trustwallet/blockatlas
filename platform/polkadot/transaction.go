package polkadot

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/txtype"
)

var NetworkByteMap = map[string]byte{
	"DOT": 0x00,
	"KSM": 0x02,
}

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	transfers, err := p.client.GetTransfersOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := make([]txtype.Tx, 0)
	for _, srcTx := range transfers {
		tx := p.NormalizeTransfer(&srcTx)
		txs = append(txs, tx)
	}

	return txs, nil
}

func (p *Platform) NormalizeTransfer(srcTx *Transfer) txtype.Tx {
	decimals := p.Coin().Decimals
	amount := strings.Split(numbers.DecimalExp(srcTx.Amount, int(decimals)), ".")[0]
	status := txtype.StatusCompleted
	if !srcTx.Success {
		status = txtype.StatusError
	}
	result := txtype.Tx{
		ID:     srcTx.Hash,
		Coin:   p.Coin().ID,
		Date:   int64(srcTx.Timestamp),
		From:   srcTx.From,
		To:     srcTx.To,
		Fee:    txtype.Amount(FeeTransfer), // API will return fee later
		Block:  srcTx.BlockNumber,
		Status: status,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(amount),
			Symbol:   p.Coin().Symbol,
			Decimals: decimals,
		},
	}
	return result
}

func (p *Platform) NormalizeExtrinsics(extrinsics []Extrinsic) []txtype.Tx {
	txs := make([]txtype.Tx, 0)
	for _, srcTx := range extrinsics {
		tx := p.NormalizeExtrinsic(&srcTx)
		if tx != nil {
			txs = append(txs, *tx)
		}
	}
	return txs
}

func (p *Platform) NormalizeExtrinsic(srcTx *Extrinsic) *txtype.Tx {
	var datas []CallData
	err := json.Unmarshal([]byte(srcTx.Params), &datas)
	if err != nil {
		return nil
	}

	// only supports balances::transfer
	if srcTx.CallModule != ModuleBalances || srcTx.CallModuleFunction != ModuleFunctionTransfer {
		return nil
	}

	// check data types
	if len(datas) < 2 || datas[0].Type != "Address" || datas[1].Type != "Compact<Balance>" {
		return nil
	}

	to := p.NormalizeAddress(datas[0].Value)
	if len(to) == 0 {
		return nil
	}

	status := txtype.StatusCompleted
	if !srcTx.Success {
		status = txtype.StatusError
	}

	result := txtype.Tx{
		ID:       srcTx.Hash,
		Coin:     p.Coin().ID,
		From:     srcTx.AccountId,
		To:       to,
		Fee:      txtype.Amount(srcTx.Fee),
		Date:     int64(srcTx.Timestamp),
		Block:    srcTx.BlockNumber,
		Status:   status,
		Sequence: srcTx.Nonce,

		Meta: txtype.Transfer{
			Value:    txtype.Amount(datas[1].Value),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}

	return &result
}

func (p *Platform) NormalizeAddress(valueRaw string) string {
	bytes, err := hex.DecodeString(valueRaw)
	if err != nil {
		return ""
	}
	if network, ok := NetworkByteMap[p.Coin().Symbol]; ok && len(bytes) > 0 {
		return PublicKeyToAddress(bytes[:], network)
	}
	return ""
}

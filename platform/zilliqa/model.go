package zilliqa

import (
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/trustwallet/blockatlas/platform/zilliqa/rpc"
	"github.com/trustwallet/blockatlas/platform/zilliqa/viewblock"
)

func TxFromRpc(t rpc.Tx, header rpc.BlockHeader) *viewblock.Tx {
	// t.recipient is not parsed correctly. Empty strings.

	to, err := hex.DecodeString(t.ToAddr)
	if err != nil {
		return nil
	}

	timestamp, err := strconv.ParseUint(header.Timestamp, 10, 64)
	if err != nil {
		timestamp = 0
	}

	height, err := strconv.ParseUint(header.Number, 10, 64)
	if err != nil {
		height = 0
	}

	gasLimit, ok := new(big.Int).SetString(t.GasLimit, 10)
	if !ok {
		return nil
	}
	gasPrice, ok := new(big.Int).SetString(t.GasPrice, 10)
	if !ok {
		return nil
	}
	fee := new(big.Int).Mul(gasLimit, gasPrice)

	return &viewblock.Tx{
		Hash:           "0x" + t.ID,
		BlockHeight:    height,
		From:           EncodePublicKeyToAddress(t.SenderPubKey),
		To:             EncodeKeyHashToAddress(to),
		Value:          t.Amount,
		Fee:            fee.String(),
		Timestamp:      int64(timestamp / 1000),
		Signature:      t.Signature,
		Nonce:          t.Nonce,
		ReceiptSuccess: t.Receipt.Success,
	}
}

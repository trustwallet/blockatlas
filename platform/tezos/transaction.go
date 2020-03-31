package tezos

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	txs, err := p.client.GetTxsOfAddress(address, txTypes)
	if err != nil {
		logger.Error("GetAddrTxs", err, logger.Params{"txType": txTypes, "addr": address})
		return nil, err
	}

	return NormalizeTxs(txs.Transactions, address), nil
}

func NormalizeTxs(srcTxs []Transaction, address string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, address)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

// NormalizeTx converts a Tezos transaction into the generic model
func NormalizeTx(srcTx Transaction, address string) (blockatlas.Tx, bool) {
	var tx = blockatlas.Tx{
		Block:  srcTx.Height,
		Coin:   coin.XTZ,
		Date:   srcTx.BlockTimestamp(),
		Error:  srcTx.ErrorMsg(),
		Fee:    blockatlas.Amount(numbers.DecimalExp(numbers.Float64toString(srcTx.Fee), 6)),
		From:   srcTx.Sender,
		ID:     srcTx.Hash,
		Status: srcTx.Status(),
		To:     srcTx.GetReceiver(),
		Type:   srcTx.TransferType(),
	}
	if address != "" {
		tx.Direction = srcTx.Direction(address)
	}

	value := blockatlas.Amount(numbers.DecimalExp(numbers.Float64toString(srcTx.Volume), 6))
	switch srcTx.TransferType() {
	case blockatlas.TxAnyAction:
		tx.Meta = blockatlas.AnyAction{
			Coin:     coin.Tezos().ID,
			Title:    srcTx.Title(),
			Key:      blockatlas.KeyStakeDelegate,
			Name:     coin.Tezos().Name,
			Symbol:   coin.Tezos().Symbol,
			Decimals: coin.Tezos().Decimals,
			Value:    value,
		}
	case blockatlas.TxTransfer:
		tx.Meta = blockatlas.Transfer{
			Value:    value,
			Symbol:   coin.Coins[coin.XTZ].Symbol,
			Decimals: coin.Coins[coin.XTZ].Decimals,
		}
	default:
		return blockatlas.Tx{}, false
	}
	return tx, true
}

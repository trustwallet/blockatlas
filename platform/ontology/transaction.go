package ontology

import (
	"errors"
	"sync"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	return p.GetTokenTxsByAddress(address, string(AssetONT))
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (types.Txs, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return types.Txs{}, err
	}
	txPage := normalizeTxs(srcTxs.Result, AssetType(token))
	return txPage, nil
}

func Normalize(srcTx *Tx, assetName AssetType) (tx types.Tx, ok bool) {
	if len(srcTx.getTransfers()) < 1 {
		return tx, false
	}
	fee := numbers.DecimalExp(srcTx.Fee, ONGDecimals)
	status := types.StatusCompleted
	if srcTx.ConfirmFlag != 1 {
		status = types.StatusError
	}
	tx = types.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.ONTOLOGY,
		Fee:    types.Amount(fee),
		Date:   srcTx.Time,
		Block:  srcTx.Height,
		Status: status,
		Memo:   "",
	}

	switch assetName {
	case AssetONT:
		return normalizeONT(tx, srcTx.getTransfers())
	case AssetONG:
		return normalizeONG(tx, srcTx.getTransfers())
	}
	return tx, false
}

func normalizeONT(tx types.Tx, transfers Transfers) (types.Tx, bool) {
	transfer := transfers.getTransfer(AssetONT)
	if transfer == nil {
		return tx, false
	}
	tx.From = transfer.FromAddress
	tx.To = transfer.ToAddress
	tx.Type = types.TxTransfer
	tx.Meta = types.Transfer{
		Value:    types.Amount(transfer.Amount),
		Symbol:   coin.Ontology().Symbol,
		Decimals: coin.Ontology().Decimals,
	}
	return tx, true
}

func normalizeONG(tx types.Tx, transfers Transfers) (types.Tx, bool) {
	transfer := transfers.getTransfer(AssetONG)
	if transfer == nil {
		return tx, false
	}
	from := transfer.FromAddress
	to := transfer.ToAddress
	tx.From = from
	tx.To = to
	value := numbers.DecimalExp(transfer.Amount, ONGDecimals)
	if transfers.isClaimReward() {
		tx.Type = types.TxAnyAction
		tx.Meta = types.AnyAction{
			Coin:     coin.Ontology().ID,
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  string(AssetONG),
			Decimals: ONGDecimals,
			Value:    types.Amount(value),
			Title:    types.AnyActionClaimRewards,
			Key:      types.KeyStakeClaimRewards,
		}
		return tx, true
	}
	tx.Type = types.TxNativeTokenTransfer
	tx.Meta = types.NativeTokenTransfer{
		Name:     "Ontology Gas",
		Symbol:   "ONG",
		TokenID:  string(AssetONG),
		Decimals: ONGDecimals,
		Value:    types.Amount(value),
		From:     from,
		To:       to,
	}
	return tx, true
}

func (p *Platform) getTxDetails(srcTx []Tx) ([]Tx, error) {
	var wg sync.WaitGroup
	txsOntV2Chan := make(chan Tx, len(srcTx))
	wg.Add(len(srcTx))
	for _, blockTxRaw := range srcTx {
		go func(blockTxRaw Tx, wg *sync.WaitGroup) {
			defer wg.Done()
			txRaw, err := p.client.GetTxDetailsByHash(blockTxRaw.Hash)
			if err == nil {
				txsOntV2Chan <- txRaw
			}
		}(blockTxRaw, &wg)
	}
	wg.Wait()
	close(txsOntV2Chan)
	if len(txsOntV2Chan) != len(srcTx) {
		return nil, errors.New("getTxDetails failed to call client.GetTxDetailsByHash http get or unmarshal")
	}
	var txsOntV2 []Tx
	for tx := range txsOntV2Chan {
		if len(tx.getTransfers()) > 0 {
			txsOntV2 = append(txsOntV2, tx)
		}
	}
	return txsOntV2, nil
}

func normalizeTxs(srcTxs []Tx, assetType AssetType) types.Txs {
	txs := make(types.Txs, 0)
	for _, srcTx := range srcTxs {
		transfer := srcTx.getTransfers().getTransfer(assetType)
		if transfer == nil {
			continue
		}
		tx, ok := Normalize(&srcTx, transfer.AssetName)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

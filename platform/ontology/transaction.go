package ontology

import (
	"github.com/trustwallet/blockatlas/coin"
	blockatlas "github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"sync"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.GetTokenTxsByAddress(address, string(AssetONT))
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		logger.Error(err, "Ontology: Failed to get transactions for address and token",
			logger.Params{
				"address": address,
				"token":   token,
			})
		return blockatlas.TxPage{}, err
	}
	txPage := normalizeTxs(srcTxs.Result, AssetType(token))
	return txPage, nil
}

func Normalize(srcTx *Tx, assetName AssetType) (tx blockatlas.Tx, ok bool) {
	if len(srcTx.getTransfers()) < 1 {
		return tx, false
	}
	fee := numbers.DecimalExp(srcTx.Fee, ONGDecimals)
	status := blockatlas.StatusCompleted
	if srcTx.ConfirmFlag != 1 {
		status = blockatlas.StatusError
	}
	tx = blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.ONT,
		Fee:    blockatlas.Amount(fee),
		Date:   srcTx.Time,
		Block:  srcTx.Height,
		Status: status,
	}

	switch assetName {
	case AssetONT:
		return normalizeONT(tx, srcTx.getTransfers())
	case AssetONG:
		return normalizeONG(tx, srcTx.getTransfers())
	}
	return tx, false
}

func normalizeONT(tx blockatlas.Tx, transfers Transfers) (blockatlas.Tx, bool) {
	transfer := transfers.getTransfer(AssetONT)
	if transfer == nil {
		return tx, false
	}
	tx.From = transfer.FromAddress
	tx.To = transfer.ToAddress
	tx.Type = blockatlas.TxTransfer
	tx.Meta = blockatlas.Transfer{
		Value:    blockatlas.Amount(transfer.Amount),
		Symbol:   coin.Ontology().Symbol,
		Decimals: coin.Ontology().Decimals,
	}
	return tx, true
}

func normalizeONG(tx blockatlas.Tx, transfers Transfers) (blockatlas.Tx, bool) {
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
		tx.Type = blockatlas.TxAnyAction
		tx.Meta = blockatlas.AnyAction{
			Coin:     coin.Ontology().ID,
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  string(AssetONG),
			Decimals: ONGDecimals,
			Value:    blockatlas.Amount(value),
			Title:    blockatlas.AnyActionClaimRewards,
			Key:      blockatlas.KeyStakeClaimRewards,
		}
		return tx, true
	}
	tx.Type = blockatlas.TxNativeTokenTransfer
	tx.Meta = blockatlas.NativeTokenTransfer{
		Name:     "Ontology Gas",
		Symbol:   "ONG",
		TokenID:  string(AssetONG),
		Decimals: ONGDecimals,
		Value:    blockatlas.Amount(value),
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
		return nil, errors.E("getTxDetails failed to call client.GetTxDetailsByHash http get or unmarshal")
	}
	var txsOntV2 []Tx
	for tx := range txsOntV2Chan {
		if len(tx.getTransfers()) > 0 {
			txsOntV2 = append(txsOntV2, tx)
		}
	}
	return txsOntV2, nil
}

func normalizeTxs(srcTxs []Tx, assetType AssetType) blockatlas.TxPage {
	var txs blockatlas.TxPage
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

package ontology

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	blockatlas "github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strings"
	"sync"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("ontology.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ONT]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.GetTokenTxsByAddress(address, string(AssetONT))
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	txPage, err := p.client.GetTxsOfAddress(address, AssetType(token))
	if err != nil {
		logger.Error(err, "Ontology: Failed to get transactions for address and token",
			logger.Params{
				"address": address,
				"token":   token,
			})
		return blockatlas.TxPage{}, err
	}
	var txs []blockatlas.Tx
	for _, srcTx := range txPage.Result.TxnList {
		tx, ok := Normalize(&srcTx, AssetType(token))
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func Normalize(srcTx *Tx, assetName AssetType) (tx blockatlas.Tx, ok bool) {
	if len(srcTx.TransferList) < 1 {
		return tx, false
	}
	fee := numbers.DecimalExp(srcTx.Fee, int(coin.Ontology().Decimals))
	status := blockatlas.StatusCompleted
	if srcTx.ConfirmFlag != 1 {
		status = blockatlas.StatusFailed
	}
	tx = blockatlas.Tx{
		ID:     srcTx.TxnHash,
		Coin:   coin.ONT,
		Fee:    blockatlas.Amount(fee),
		Date:   srcTx.TxnTime,
		Block:  srcTx.Height,
		Status: status,
	}

	switch assetName {
	case AssetONT:
		return normalizeONT(tx, srcTx.TransferList)
	case AssetONG:
		return normalizeONG(tx, srcTx.TransferList)
	}
	return tx, false
}

func normalizeONT(tx blockatlas.Tx, transfers Transfers) (blockatlas.Tx, bool) {
	transfer := transfers.getTransfer()
	if transfer == nil {
		return tx, false
	}
	i := strings.IndexRune(transfer.Amount, '.')
	value := transfer.Amount[:i]
	tx.From = transfer.FromAddress
	tx.To = transfer.ToAddress
	tx.Type = blockatlas.TxTransfer
	tx.Meta = blockatlas.Transfer{
		Value:    blockatlas.Amount(value),
		Symbol:   coin.Ontology().Symbol,
		Decimals: coin.Ontology().Decimals,
	}
	return tx, true
}

func normalizeONG(tx blockatlas.Tx, transfers Transfers) (blockatlas.Tx, bool) {
	transfer := transfers.getTransfer()
	if transfer == nil {
		return tx, false
	}

	from := transfer.FromAddress
	to := transfer.ToAddress
	tx.From = from
	tx.To = to
	tx.Type = blockatlas.TxNativeTokenTransfer
	decimals := coin.Ontology().Decimals
	value := numbers.DecimalExp(transfer.Amount, int(decimals))
	if transfers.isClaimReward() {
		tx.Meta = blockatlas.AnyAction{
			Coin:     coin.Ontology().ID,
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  string(AssetONG),
			Decimals: decimals,
			Value:    blockatlas.Amount(value),
			Title:    blockatlas.AnyActionClaimRewards,
			Key:      blockatlas.KeyStakeClaimRewards,
		}
		return tx, true
	}
	tx.Meta = blockatlas.NativeTokenTransfer{
		Name:     "Ontology Gas",
		Symbol:   "ONG",
		TokenID:  string(AssetONG),
		Decimals: decimals,
		Value:    blockatlas.Amount(value),
		From:     from,
		To:       to,
	}
	return tx, true
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.CurrentBlockNumber()
	if err != nil {
		return 0, errors.E(err, "CurrentBlockNumber")
	}
	if len(block.Result) == 0 {
		return 0, errors.E("invalid block height result")
	}
	return block.Result[0].Height, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	blockOnt, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txsRaw, err := p.getTxDetails(blockOnt.Result.TxnList)
	if err != nil {
		return nil, err
	}
	block := normalizeBlock(*blockOnt, txsRaw)
	return block, nil
}

func (p *Platform) getTxDetails(txsOntV1 []Tx) ([]TxV2, error) {
	var wg sync.WaitGroup
	txsOntV2Chan := make(chan TxV2, len(txsOntV1))
	wg.Add(len(txsOntV1))
	for _, blockTxRaw := range txsOntV1 {
		go func(blockTxRaw Tx, wg *sync.WaitGroup) {
			defer wg.Done()
			txRaw, err := p.client.GetTxDetailsByHash(blockTxRaw.TxnHash)
			if err == nil {
				txsOntV2Chan <- *txRaw
			}
		}(blockTxRaw, &wg)
	}
	wg.Wait()
	close(txsOntV2Chan)
	if len(txsOntV2Chan) != len(txsOntV1) {
		return nil, errors.E("getTxDetails failed to call client.GetTxDetailsByHash http get or unmarshal")
	}
	var txsOntV2 []TxV2
	for tx := range txsOntV2Chan {
		if len(tx.Details.Transfers) > 0 {
			txsOntV2 = append(txsOntV2, tx)
		}
	}
	return txsOntV2, nil
}

func normalizeBlock(blockOnt BlockResult, txsOntV2 []TxV2) *blockatlas.Block {
	block := blockatlas.Block{
		Number: blockOnt.Result.Height,
		ID:     blockOnt.Result.Hash,
	}
	if len(txsOntV2) > 0 {
		block.Txs = normalizeBlockTransactions(txsOntV2)
	}
	return &block
}

func normalizeBlockTransactions(txsOntV2 []TxV2) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, txOntV2 := range txsOntV2 {
		tx, ok := normalizeBlockTransaction(txOntV2)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

func normalizeBlockTransaction(txsOntV2 TxV2) (blockatlas.Tx, bool) {
	if len(txsOntV2.Details.Transfers) < 1 {
		return blockatlas.Tx{}, false
	}
	fee := numbers.DecimalExp(txsOntV2.Fee, int(coin.Ontology().Decimals))
	status := blockatlas.StatusCompleted
	if txsOntV2.ConfirmFlag != 1 {
		status = blockatlas.StatusFailed
	}
	tx := blockatlas.Tx{
		ID:     txsOntV2.Hash,
		Coin:   coin.ONT,
		Fee:    blockatlas.Amount(fee),
		Date:   txsOntV2.Time,
		Block:  txsOntV2.BlockHeight,
		Status: status,
	}
	transfer := txsOntV2.Details.Transfers.getTransfer()
	if transfer == nil {
		return tx, false
	}
	switch transfer.AssetName {
	case AssetONT:
		return normalizeONT(tx, txsOntV2.Details.Transfers)
	case AssetONG:
		return normalizeONG(tx, txsOntV2.Details.Transfers)
	}
	return tx, false
}

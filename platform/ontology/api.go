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

const (
	GovernanceContract = "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK"
	ONTAssetName       = "ont"
	ONGAssetName       = "ong"
	ONGDecimals        = 9
)

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("ontology.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ONT]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.GetTokenTxsByAddress(address, ONTAssetName)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	txPage, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		logger.Error(err, "Ontology: Failed to get transactions for address and token",
			logger.Params{
				"address": address,
				"token":   token,
			})
		return blockatlas.TxPage{}, err
	}
	var txs []blockatlas.Tx
	for _, txOntV1 := range txPage.Result.TxnList {
		tx, ok := Normalize(txOntV1, token)
		if !ok {
			continue
		}
		txs = append(txs, *tx)
	}

	return txs, nil
}

func Normalize(txOntV1 Tx, assetName string) (tx *blockatlas.Tx, ok bool) {
	if len(txOntV1.TransferList) < 1 {
		return tx, false
	}
	transfer := txOntV1.TransferList[0]
	fee := numbers.DecimalExp(txOntV1.Fee, ONGDecimals)
	tx = &blockatlas.Tx{
		ID:     txOntV1.TxnHash,
		Coin:   coin.ONT,
		Date:   txOntV1.TxnTime,
		Block:  txOntV1.Height,
		Status: getTransactionStatus(int(txOntV1.ConfirmFlag)),
		Fee:    blockatlas.Amount(fee),
	}

	switch assetName {
	case ONTAssetName:
		normalizeONT(tx, transfer)
	case ONGAssetName:
		normalizeONG(tx, transfer)
	default: // unsupported asset
		return tx, false
	}
	return tx, true
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.CurrentBlockNumber()
	if err != nil {
		logger.Error("CurrentBlockNumber", logger.Params{"platform": p.Coin().Symbol, "details": err.Error()})
		return 0, err
	}
	var height int64
	if len(block.Result) > 0 {
		height = (int64)(block.Result[0].Height)
	}
	return height, nil
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
		Number: int64(blockOnt.Result.Height),
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
		txs = append(txs, *tx)
	}
	return txs
}

func normalizeBlockTransaction(txsOntV2 TxV2) (*blockatlas.Tx, bool) {
	fee := numbers.DecimalExp(txsOntV2.Fee, ONGDecimals)
	tx := blockatlas.Tx{
		ID:     txsOntV2.Hash,
		Coin:   coin.ONT,
		Date:   txsOntV2.Time,
		Block:  txsOntV2.BlockHeight,
		Status: getTransactionStatus(txsOntV2.ConfirmFlag),
		Fee:    blockatlas.Amount(fee),
	}
	ok := false
	if len(txsOntV2.Details.Transfers) > 0 {
		txDetails := txsOntV2.Details.Transfers[0]
		ok = true
		switch txDetails.AssetName {
		case ONTAssetName:
			normalizeONT(&tx, Transfer{
				Amount:      txDetails.Amount,
				FromAddress: txDetails.FromAddress,
				ToAddress:   txDetails.ToAddress,
			})
		case ONGAssetName:
			normalizeONG(&tx, Transfer{
				Amount:      txDetails.Amount,
				FromAddress: txDetails.FromAddress,
				ToAddress:   txDetails.ToAddress,
			})
		default: // unsupported asset
			return &tx, false
		}
	}
	return &tx, ok
}

func getTransactionStatus(confirmFlag int) blockatlas.Status {
	if confirmFlag == 1 {
		return blockatlas.StatusCompleted
	}
	return blockatlas.StatusFailed
}

func normalizeONT(tx *blockatlas.Tx, transfer Transfer) {
	i := strings.IndexRune(transfer.Amount, '.')
	var value string
	if i > 0 {
		value = transfer.Amount[:i]
	} else {
		value = transfer.Amount
	}

	tx.From = transfer.FromAddress
	tx.To = transfer.ToAddress
	tx.Type = blockatlas.TxTransfer
	tx.Meta = blockatlas.Transfer{
		Value:    blockatlas.Amount(value),
		Symbol:   coin.Coins[coin.ONT].Symbol,
		Decimals: coin.Coins[coin.ONT].Decimals,
	}
}

func normalizeONG(tx *blockatlas.Tx, transfer Transfer) {
	var value string
	if transfer.ToAddress == GovernanceContract {
		value = "0"
	} else {
		value = numbers.DecimalExp(transfer.Amount, ONGDecimals)
	}

	from := transfer.FromAddress
	to := transfer.ToAddress
	tx.From = from
	tx.To = to
	tx.Type = blockatlas.TxNativeTokenTransfer
	tx.Meta = blockatlas.NativeTokenTransfer{
		Name:     "Ontology Gas",
		Symbol:   "ONG",
		TokenID:  "ong",
		Decimals: ONGDecimals,
		Value:    blockatlas.Amount(value),
		From:     from,
		To:       to,
	}
}

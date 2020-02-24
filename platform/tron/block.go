package tron

import (
	"encoding/hex"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sync"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	block, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txsChan := p.NormalizeBlockTxs(block.Txs)
	txs := make(blockatlas.TxPage, 0)
	for cTxs := range txsChan {
		txs = append(txs, cTxs)
	}

	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) NormalizeBlockTxs(srcTxs []Tx) chan blockatlas.Tx {
	txChan := make(chan blockatlas.Tx, len(srcTxs))
	var wg sync.WaitGroup
	for _, srcTx := range srcTxs {
		wg.Add(1)
		go func(s Tx, c chan blockatlas.Tx) {
			defer wg.Done()
			p.NormalizeBlockChannel(s, c)
		}(srcTx, txChan)
	}
	wg.Wait()
	close(txChan)
	return txChan
}

func (p *Platform) NormalizeBlockChannel(srcTx Tx, txChan chan blockatlas.Tx) {
	if len(srcTx.Data.Contracts) == 0 {
		return
	}

	tx, err := Normalize(srcTx)
	if err != nil {
		return
	}
	transfer := srcTx.Data.Contracts[0].Parameter.Value
	if len(transfer.AssetName) > 0 {
		assetName, err := hex.DecodeString(transfer.AssetName[:])
		if err == nil {
			info, err := p.client.GetTokenInfo(string(assetName))
			if err == nil && len(info.Data) > 0 {
				setTokenMeta(tx, srcTx, info.Data[0])
			}
		}
	}
	txChan <- *tx
}

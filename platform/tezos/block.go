package tezos

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	srcTxs, err := p.client.GetBlockByNumber(num, txTypes)
	if err != nil {
		log.WithFields(log.Fields{"txType": txTypes, "num": num}).Error("GetAddrTxs", err)
		return nil, err
	}
	txs := NormalizeTxs(srcTxs, "")
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

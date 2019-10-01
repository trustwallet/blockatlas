package waves

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"

	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("waves.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.WAVES]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	currentBlock, err := p.client.GetCurrentBlock()
	if err != nil {
		return 0, err
	}
	return currentBlock.Height, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcTxs.Transactions)

	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, 25)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(addressTxs)

	return txs, nil
}

func NormalizeTxs(srcTxs []Transaction) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

func NormalizeTx(srcTx *Transaction) (tx blockatlas.Tx, ok bool) {
	var result blockatlas.Tx

	if srcTx.Type == 4 && len(srcTx.AssetId) == 0 {
		result = blockatlas.Tx{
			ID:     srcTx.Id,
			Coin:   coin.WAVES,
			From:   srcTx.Sender,
			To:     srcTx.Recipient,
			Fee:    blockatlas.Amount(strconv.Itoa(int(srcTx.Fee))),
			Date:   int64(srcTx.Timestamp) / 1000,
			Block:  srcTx.Block,
			Memo:   srcTx.Attachment,
			Status: blockatlas.StatusCompleted,
			Meta: blockatlas.Transfer{
				Value:    blockatlas.Amount(strconv.Itoa(int(srcTx.Amount))),
				Symbol:   coin.Coins[coin.WAVES].Symbol,
				Decimals: coin.Coins[coin.WAVES].Decimals,
			},
		}
		return result, true
	}

	return result, false
}

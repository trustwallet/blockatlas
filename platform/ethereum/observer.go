package ethereum

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/platform/ethereum/source"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
	"time"
)

var dispatcher *observer.Dispatcher
var interval time.Duration
var client *source.Client

func SetupObserver(d *observer.Dispatcher, sleepInterval time.Duration) {
	dispatcher = d
	interval = sleepInterval
	client = source.NewClient(viper.GetString("ethereum.api"))
}

func ObserveNewBlocks() {
	if dispatcher == nil {
		logrus.Error("Please, run SetupObserver function before start listening")
		return
	}

	ethApi := viper.GetString("ethereum.api")
	logrus.Infof("ETH: Observing new blocks from %s each %d seconds", ethApi, interval)

	var currentBlockNumber uint64

	observeLoop(func(blockNumber uint64) {
		if currentBlockNumber == 0 {
			currentBlockNumber = blockNumber
			go dispatchBlock(currentBlockNumber)
			return
		}

		if blockNumber > currentBlockNumber {
			var n, diff uint64; n = 1; diff = blockNumber - currentBlockNumber
			for ; n <= diff; n++ {
				go dispatchBlock(currentBlockNumber + n)
			}
			currentBlockNumber = blockNumber
		}
	})
}

func observeLoop(f func(uint64)) {
	for {
		if blockNumber, err := client.BlockNumber(); err == nil && blockNumber > 0 {
			f(blockNumber)
		} else {
			logrus.WithError(err).Error("Failed to get latest block")
		}

		time.Sleep(interval * time.Second)
	}
}

func dispatchBlock(blockNumber uint64) {
	block, err := client.GetBlockByNumber(blockNumber)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to fetch block n %d", blockNumber)
		return
	}

	txs := make([]models.Tx, 0)
	for _, srcTx := range block.Transactions {
		if len(srcTx.To) == 0 {
			continue
		}

		txs = append(txs, models.Tx{
			Coin: coin.IndexETH,
			Type: models.TxTransfer,
			Id: srcTx.Hash,
			From: srcTx.From,
			To: srcTx.To,
			Date: int64(block.Timestamp),
			Block: uint64(block.Number),
			Fee: strconv.FormatUint(uint64(srcTx.Gas * srcTx.GasPrice), 10),
			Meta: models.Transfer{
				Value: util.DecimalExp(srcTx.Value.ToInt().String(), int(coin.ETH.Decimals)),
			},
		})
	}

	dispatcher.DispatchTransactions(txs)
}

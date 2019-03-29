package ethereum

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/observer"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var dispatcher *observer.Dispatcher
var queue = make([]*types.Header, 0)
var minBlockDelay = 3

func Setup(d *observer.Dispatcher, delay int) {
	dispatcher = d
	minBlockDelay = delay
}

func ObserveNewBlocks() {
	if dispatcher == nil {
		logrus.Error("Please, run Setup function before start listening")
		return
	}

	ws := viper.GetString("ethereum.ws")
	client, err := ethclient.Dial(ws)

	if err != nil {
		logrus.WithError(err).Error("Failed to connect to endpoint")
		return
	}

	logrus.Infof("ETH: Observing new blocks from %s", ws)

	newHeaders := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), newHeaders)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to subscribe")
		return
	}

	for {
		select {
		case err := <- sub.Err():
			logrus.WithError(err)
		case header := <-newHeaders:
			enqueue(client, header)
		}
	}
}

func enqueue(client *ethclient.Client, header *types.Header, ) {
	logrus.Debugf("Enqueueing header %s", header.Hash().String())

	queue = append(queue, header)

	if len(queue) > minBlockDelay {
		var h *types.Header
		h, queue = queue[0], queue[1:] // Pop the first header in queue
		go process(client, h)
	}
}

func process(client *ethclient.Client, header *types.Header) {
	chainID := new(big.Int).SetInt64(viper.GetInt64("ethereum.chainID"))
	block, err := client.BlockByHash(context.Background(), header.Hash())

	if err != nil {
		logrus.WithError(err).Error("Failed to get block")
		return
	}

	logrus.Debugf("Processing block %s", block.Hash().String())

	var txs []models.Tx
	for _, blockTx := range block.Transactions() {
		if msg, err := blockTx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			if msg.To() == nil {
				continue
			}

			tx := models.Tx{
				Id: blockTx.Hash().String(),
				Coin: coin.ETH.Index,
				To: msg.To().Hex(),
				From: msg.From().Hex(),
				Fee: blockTx.Cost().String(),
				Block: block.NumberU64(),
				Date: block.Time().Int64(),
				Type: models.TxTransfer,
				Meta: models.Transfer{
					Value:blockTx.Value().String(),
				},
			}

			txs = append(txs, tx)
		}
	}

	dispatcher.DispatchTransactions(txs)
}

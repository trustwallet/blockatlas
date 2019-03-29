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

func ListenForLatestBlock(dispatcher observer.Dispatcher) {
	ws := viper.GetString("ethereum.ws")
	chainID := new(big.Int).SetInt64(viper.GetInt64("ethereum.chainID"))

	client, err := ethclient.Dial(ws)
	logrus.Infof("ETH: Observing new blocks from %s", ws)

	if err != nil {
		logrus.WithError(err).Error("Failed to connect to endpoint")
		return
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to subscribe")
		return
	}

	for {
		select {
		case err := <-sub.Err():
			logrus.WithError(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				logrus.WithError(err).Error("Failed to get block")
				break
			}

			logrus.Debugf("Block %s", block.Hash().String())

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

			dispatcher.NotifyObservers(txs)
		}
	}
}

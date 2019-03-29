package ethereum

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/observer"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func StartObserver(dispatcher observer.Dispatcher) {
	chainID := new(big.Int).SetInt64(viper.GetInt64("ethereum.chainID"))
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")

	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			var txs []models.Tx
			for _, blockTx := range block.Transactions() {
				if msg, err := blockTx.AsMessage(types.NewEIP155Signer(chainID)); err != nil {
					tx := models.Tx{
						Id: header.Hash().Hex(),
						Coin: coin.ETH.Index,
						To: msg.To().Hex(),
						From: msg.From().Hex(),
						Fee: blockTx.Cost().String(),
						Block: block.NumberU64(),
						Date: block.Time().Int64(),
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

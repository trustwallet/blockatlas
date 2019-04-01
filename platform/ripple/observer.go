package ripple

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/observer"
	"github.com/trustwallet/blockatlas/platform/ripple/source"
	"github.com/trustwallet/blockatlas/util"
	"github.com/valyala/fastjson"
	"strconv"
	"time"
)

var dispatcher *observer.Dispatcher

func SetupObserver(d *observer.Dispatcher) {
	dispatcher = d
	client.WsUrl = viper.GetString("ripple.ws")
}

func ObserveNewBlocs() {
	if dispatcher == nil {
		logrus.Error("Please, run SetupObserver function before start listening")
		return
	}

	cError := make(chan error)
	cLedger := make(chan source.Ledger)

	err := client.SubscribeLedger(cLedger, cError)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	logrus.Infof("XRP: Observing new blocks from %s", client.WsUrl)

	for {
		select {
		case err := <-cError:
			logrus.Error(err)
		case ledger := <-cLedger:
			date, err := time.Parse("2006-Jan-02 15:04:05", ledger.CloseTimeHuman)
			legerClose := date.Unix()
			if err != nil {
				legerClose = 0
			}

			ledgerIndex, err := strconv.ParseInt(ledger.Index, 10, 64)
			if err != nil {
				ledgerIndex = 0
			}

			txs := make([]models.Tx, 0)
			for _, srcTx := range ledger.Transactions {
				if srcTx.TransactionType != "Payment" {
					continue
				}
				// Only accept XRP payments (typeof tx.amount === 'string')
				var p fastjson.Parser
				v, pErr := p.ParseBytes(srcTx.Amount)
				if pErr != nil {
					continue
				}
				if v.Type() != fastjson.TypeString {
					continue
				}

				srcAmount := string(v.GetStringBytes())
				txs = append(txs, models.Tx{
					Id:    srcTx.Hash,
					Coin:  coin.IndexXRP,
					Date:  legerClose,
					From:  srcTx.Account,
					To:    srcTx.Destination,
					Fee:   util.DecimalExp(srcTx.Fee, 6),
					Block: uint64(ledgerIndex),
					Meta: models.Transfer{
						Value: util.DecimalExp(srcAmount, 6),
					},
				})
			}

			dispatcher.DispatchTransactions(txs)
		}
	}
}

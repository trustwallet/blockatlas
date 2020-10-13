package tokenindexer

import (
	"context"

	"github.com/streadway/amqp"
	"go.elastic.co/apm"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/notifier"
)

func RunTokenIndexer(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunTokenIndexer", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	txs, err := notifier.GetTransactionsFromDelivery(delivery, ctx)
	if err != nil {
		logger.Error("failed to get transactions", err)
		return
	}
	if len(txs) == 0 {
		return
	}

	types := getTokenTypes(txs)
	if len(types) == 0 {
		return
	}

	if err = database.CreateTokenTypes(ctx, types); err != nil {
		logger.Error("failed to create token types", err)
	}
}

func getTokenTypes(txs blockatlas.Txs) []string {
	var types []string
	for _, tx := range txs {
		var tokenType string
		switch tx.Type {
		case blockatlas.TxTokenTransfer:
			tokenMeta, ok := tx.Meta.(*blockatlas.TokenTransfer)
			if !ok {
				continue
			}

			tokenType, ok = blockatlas.GetTokenType(tokenMeta.Symbol)
			if !ok {
				continue
			}

		default:
			continue
		}
		types = append(types, tokenType)
	}
	return types
}

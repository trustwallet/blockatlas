package tokenindexer

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.elastic.co/apm"
)

type Params struct {
	Ctx      context.Context
	Queue    mq.Queue
	Database *db.Instance
}

func RunTokenIndexer(params Params) {
	tx := apm.DefaultTracer.StartTransaction("RunTokenIndexer", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	for v := range params.Queue.GetMessageChannel() {
		if err := getOrCreateTokenTypes(ctx, params.Database, v); err != nil {
			logger.Warn("tokenindexer.getOrCreateTokenTypes: err: ", err)
		}
	}
}

func getOrCreateTokenTypes(ctx context.Context, db *db.Instance, v amqp.Delivery) error {
	parsedTokenTypes, err := parseMessage(v.Body)
	if err != nil {
		return err
	}
	return db.CreateTokenType(ctx, parsedTokenTypes)
}

func parseMessage(body []byte) ([]models.TokenType, error) {
	var txs blockatlas.Txs
	if err := json.Unmarshal(body, &txs); err != nil {
		return nil, err
	}

	tokenTypes := make([]models.TokenType, 0, len(txs))
	for _, tx := range txs {
		if tx.Type == blockatlas.TxTokenTransfer {
			tokenMeta, ok := tx.Meta.(blockatlas.TokenTransfer)
			if !ok {
				continue
			}

			if tokenMeta.Name == "" || tokenMeta.Symbol == "" {
				continue
			}
			tokenTypes = append(tokenTypes, models.TokenType{
				Type:     tokenMeta.TokenID,
				Decimals: tokenMeta.Decimals,
				Name:     tokenMeta.Name,
				Symbol:   tokenMeta.Symbol,
			})
		}
	}
	return tokenTypes, nil
}

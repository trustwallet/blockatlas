package tokensearcher

import (
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/watchmarket/pkg/watchmarket"
	"strconv"
)

func makeTransactionsMap(coin uint, txs blockatlas.Txs) map[string]blockatlas.Txs {
	result := make(map[string]blockatlas.Txs)
	var prefix = strconv.Itoa(int(coin)) + "_"
	for _, tx := range txs {
		_, from, to, ok := getInfoByMeta(tx)
		if !ok {
			continue
		}
		result[prefix+from] = append(result[prefix+from], tx)
		result[prefix+to] = append(result[prefix+to], tx)
	}
	return result
}

func getInfoByMeta(tx blockatlas.Tx) (models.Token, string, string, bool) {
	var (
		token    models.Token
		from, to string
	)
	switch tx.Meta.(type) {
	case blockatlas.TokenTransfer:
		from = tx.Meta.(blockatlas.TokenTransfer).From
		to = tx.Meta.(blockatlas.TokenTransfer).To
		token.Coin = tx.Coin
		token.AssetID = watchmarket.BuildID(tx.Coin, tx.Meta.(blockatlas.TokenTransfer).TokenID)
		token.Address = tx.Meta.(blockatlas.TokenTransfer).TokenID
		token.TokenID = tx.Meta.(blockatlas.TokenTransfer).TokenID
		token.Decimals = int(tx.Meta.(blockatlas.TokenTransfer).Decimals)
		token.Symbol = tx.Meta.(blockatlas.TokenTransfer).Symbol
		token.Name = tx.Meta.(blockatlas.TokenTransfer).Name
	case *blockatlas.TokenTransfer:
		from = tx.Meta.(*blockatlas.TokenTransfer).From
		to = tx.Meta.(*blockatlas.TokenTransfer).To
		token.Coin = tx.Coin
		token.AssetID = watchmarket.BuildID(tx.Coin, tx.Meta.(*blockatlas.TokenTransfer).TokenID)
		token.Address = tx.Meta.(*blockatlas.TokenTransfer).TokenID
		token.TokenID = tx.Meta.(*blockatlas.TokenTransfer).TokenID
		token.Decimals = int(tx.Meta.(*blockatlas.TokenTransfer).Decimals)
		token.Symbol = tx.Meta.(*blockatlas.TokenTransfer).Symbol
		token.Name = tx.Meta.(*blockatlas.TokenTransfer).Name
	case blockatlas.NativeTokenTransfer:
		from = tx.Meta.(blockatlas.NativeTokenTransfer).From
		to = tx.Meta.(blockatlas.NativeTokenTransfer).To
		token.Coin = tx.Coin
		token.AssetID = watchmarket.BuildID(tx.Coin, tx.Meta.(blockatlas.NativeTokenTransfer).TokenID)
		token.Address = tx.Meta.(blockatlas.NativeTokenTransfer).TokenID
		token.TokenID = tx.Meta.(blockatlas.NativeTokenTransfer).TokenID
		token.Decimals = int(tx.Meta.(blockatlas.NativeTokenTransfer).Decimals)
		token.Symbol = tx.Meta.(blockatlas.NativeTokenTransfer).Symbol
		token.Name = tx.Meta.(blockatlas.NativeTokenTransfer).Name
	case *blockatlas.NativeTokenTransfer:
		from = tx.Meta.(*blockatlas.NativeTokenTransfer).From
		to = tx.Meta.(*blockatlas.NativeTokenTransfer).To
		token.Coin = tx.Coin
		token.AssetID = watchmarket.BuildID(tx.Coin, tx.Meta.(*blockatlas.NativeTokenTransfer).TokenID)
		token.Address = tx.Meta.(*blockatlas.NativeTokenTransfer).TokenID
		token.TokenID = tx.Meta.(*blockatlas.NativeTokenTransfer).TokenID
		token.Decimals = int(tx.Meta.(*blockatlas.NativeTokenTransfer).Decimals)
		token.Symbol = tx.Meta.(*blockatlas.NativeTokenTransfer).Symbol
		token.Name = tx.Meta.(*blockatlas.NativeTokenTransfer).Name
	case blockatlas.AnyAction:
		from = tx.From
		to = tx.To
		token.Coin = tx.Coin
		token.AssetID = watchmarket.BuildID(tx.Coin, tx.Meta.(blockatlas.AnyAction).TokenID)
		token.Address = tx.Meta.(blockatlas.AnyAction).TokenID
		token.TokenID = tx.Meta.(blockatlas.AnyAction).TokenID
		token.Decimals = int(tx.Meta.(blockatlas.AnyAction).Decimals)
		token.Symbol = tx.Meta.(blockatlas.AnyAction).Symbol
		token.Name = tx.Meta.(blockatlas.AnyAction).Name
	case *blockatlas.AnyAction:
		from = tx.From
		to = tx.To
		token.Coin = tx.Coin
		token.AssetID = watchmarket.BuildID(tx.Coin, tx.Meta.(*blockatlas.AnyAction).TokenID)
		token.Address = tx.Meta.(*blockatlas.AnyAction).TokenID
		token.TokenID = tx.Meta.(*blockatlas.AnyAction).TokenID
		token.Decimals = int(tx.Meta.(*blockatlas.AnyAction).Decimals)
		token.Symbol = tx.Meta.(*blockatlas.AnyAction).Symbol
		token.Name = tx.Meta.(*blockatlas.AnyAction).Name
	default:
		return models.Token{}, from, to, false
	}
	return token, from, to, true
}

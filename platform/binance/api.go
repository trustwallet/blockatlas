package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"strings"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = ClientInit(http.DefaultClient, viper.GetString("binance.api"), viper.GetString("binance.dex"))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	// No native function to get height in explorer API
	// Workaround: Request list of blocks
	// and return number of the newest one
	list, err := p.client.GetBlockList(1)
	if err != nil {
		return 0, err
	}
	if len(list.BlockArray) == 0 {
		return 0, fmt.Errorf("no block descriptor found")
	}
	return list.BlockArray[0].BlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	// TODO: Only returns BNB transactions for now
	txs := NormalizeTxs(srcTxs.Txs, "", len(srcTxs.Txs))
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// Endpoint supports queries without token query parameter
	return p.GetTokenTxsByAddress(address, "")
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(srcTxs.Txs, token, blockatlas.TxPerPage), nil
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx *Tx, token string) (tx blockatlas.Tx, ok bool) {
	value := util.DecimalExp(string(srcTx.Value), 8)
	fee := util.DecimalExp(string(srcTx.Fee), 8)

	tx = blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.BNB,
		Date:  srcTx.Timestamp / 1000,
		From:  srcTx.FromAddr,
		To:    srcTx.ToAddr,
		Fee:   blockatlas.Amount(fee),
		Block: srcTx.BlockHeight,
		Memo:  srcTx.Memo,
	}

	// Condition for native transfer (BNB)
	if srcTx.Asset == "BNB" && srcTx.Type == "TRANSFER" && token == "" {
		tx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[coin.BNB].Symbol,
			Decimals: coin.Coins[coin.BNB].Decimals,
		}
		return tx, true
	}

	// Condition for native token transfer
	if srcTx.Asset == token && srcTx.Type == "TRANSFER" && srcTx.FromAddr != "" && srcTx.ToAddr != "" {
		tx.Meta = blockatlas.NativeTokenTransfer{
			TokenID:  srcTx.Asset,
			Symbol:   TokenSymbol(srcTx.Asset),
			Value:    blockatlas.Amount(value),
			Decimals: coin.Coins[coin.BNB].Decimals,
			From:     srcTx.FromAddr,
			To:       srcTx.ToAddr,
		}

		return tx, true
	}

	return tx, false
}

func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []Tx, token string, pageSize int) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx, token)
		if !ok || len(txs) >= pageSize {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.GetAccountMetadata(address)
	if err != nil {
		return nil, err
	}
	tokens, err := p.client.GetTokens()
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Balances, tokens), nil
}

// NormalizeToken converts a Binance token into the generic model
func NormalizeToken(srcToken *Balance, tokens *TokenPage) (t blockatlas.Token, ok bool) {
	tk := tokens.findToken(srcToken.Symbol)
	if tk == nil {
		return blockatlas.Token{}, false
	}

	t = blockatlas.Token{
		Name:     tk.Name,
		Symbol:   tk.OriginalSymbol,
		TokenId:  tk.Symbol,
		Coin:     coin.BNB,
		Decimals: uint(decimalPlaces(tk.TotalSupply)),
	}

	return t, true
}

// NormalizeTxs converts multiple Binance tokens
func NormalizeTokens(srcBalance []Balance, tokens *TokenPage) (tokenPage []blockatlas.Token) {
	for _, srcToken := range srcBalance {
		token, ok := NormalizeToken(&srcToken, tokens)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return
}

// decimalPlaces count the decimals places.
func decimalPlaces(v string) int {
	s := strings.Split(v, ".")
	if len(s) < 2 {
		return 0
	}
	return len(s[1])
}

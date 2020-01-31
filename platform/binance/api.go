package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client    Client
	dexClient DexClient
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("binance.api"))}
	p.client.ErrorHandler = getHTTPError

	p.dexClient = DexClient{blockatlas.InitClient(viper.GetString("binance.dex"))}
	p.dexClient.ErrorHandler = getHTTPError
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
		return 0, errors.E("no block descriptor found", errors.TypePlatformApi).PushToSentry()
	}
	return list.BlockArray[0].BlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	childTxs, err := p.GetTxChildChan(srcTxs.Txs)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(childTxs, "", "")
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// Endpoint supports queries without token query parameter
	return p.GetTokenTxsByAddress(address, p.Coin().Symbol)
}

func (p *Platform) GetTxChildChan(srcTxs []Tx) ([]Tx, error) {
	txs := make([]Tx, 0)

	var wg sync.WaitGroup
	out := make(chan Tx)
	for _, srcTx := range srcTxs {
		if srcTx.HasChildren != 1 {
			txs = append(txs, srcTx)
			continue
		}
		wg.Add(1)
		go func(srcTx Tx, out chan Tx) {
			defer wg.Done()
			tx, err := p.client.GetTx(srcTx.Hash)
			if err != nil {
				out <- srcTx
				logger.Error("GetTransactionsByBlockChan", err, logger.Params{"hash": srcTx.Hash})
				return
			}
			out <- tx
		}(srcTx, out)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	for r := range out {
		txs = append(txs, r)
	}
	return txs, nil
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}
	txs, err := p.GetTxChildChan(srcTxs.Txs)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(txs, token, address), nil
}

func normalizeTransfer(tx blockatlas.Tx, srcTx Tx, token, address string) (blockatlas.TxPage, bool) {
	if srcTx.HasChildren == 1 {
		txs := make(blockatlas.TxPage, 0)
		for _, subTx := range srcTx.SubTxsDto.SubTxDtoList.getTxs() {
			if !subTx.containAddress(address) {
				continue
			}
			newTxs, ok := normalizeTransfer(tx, subTx, token, address)
			if !ok {
				continue
			}
			txs = append(txs, newTxs...)
		}
		return txs, true
	}

	if len(token) > 0 && srcTx.Asset != token {
		return blockatlas.TxPage{tx}, false
	}

	tx.From = srcTx.FromAddr
	tx.To = srcTx.ToAddr

	bnbCoin := coin.Coins[coin.BNB]
	value := numbers.DecimalExp(string(srcTx.Value), 8)
	if srcTx.Asset == bnbCoin.Symbol {
		// Condition for native transfer (BNB)
		tx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   bnbCoin.Symbol,
			Decimals: bnbCoin.Decimals,
		}
		return blockatlas.TxPage{tx}, true
	}

	// Condition for native token transfer
	tx.Meta = blockatlas.NativeTokenTransfer{
		TokenID:  srcTx.Asset,
		Symbol:   TokenSymbol(srcTx.Asset),
		Value:    blockatlas.Amount(value),
		Decimals: bnbCoin.Decimals,
		From:     srcTx.FromAddr,
		To:       srcTx.ToAddr,
	}
	return blockatlas.TxPage{tx}, true
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx Tx, token, address string) (blockatlas.TxPage, bool) {
	tx := blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.BNB,
		From:   srcTx.FromAddr,
		To:     srcTx.ToAddr,
		Fee:    blockatlas.Amount(srcTx.getFee()),
		Date:   srcTx.Timestamp / 1000,
		Block:  srcTx.BlockHeight,
		Status: blockatlas.StatusCompleted,
		Memo:   srcTx.Memo,
	}

	switch srcTx.Type {
	case TxTransfer:
		return normalizeTransfer(tx, srcTx, token, address)
	}
	//case TxCancelOrder, TxNewOrder:
	//	return tx, false
	//	dt, err := srcTx.getData()
	//	if err != nil {
	//		return tx, false
	//	}
	//
	//	symbol := dt.OrderData.Quote
	//	if len(token) > 0 && symbol != token {
	//		return tx, false
	//	}
	//
	//	key := blockatlas.KeyPlaceOrder
	//	title := blockatlas.KeyTitlePlaceOrder
	//	if srcTx.Type == TxCancelOrder {
	//		key = blockatlas.KeyCancelOrder
	//		title = blockatlas.KeyTitleCancelOrder
	//	}
	//	volume, ok := dt.OrderData.GetVolume()
	//	if ok {
	//		value = strconv.Itoa(int(volume))
	//	}
	//
	//	tx.Meta = blockatlas.AnyAction{
	//		Coin:     coin.BNB,
	//		TokenID:  dt.OrderData.Symbol,
	//		Symbol:   TokenSymbol(symbol),
	//		Name:     symbol,
	//		Value:    blockatlas.Amount(value),
	//		Decimals: coin.Coins[coin.BNB].Decimals,
	//		Title:    title,
	//		Key:      key,
	//	}
	//}
	return blockatlas.TxPage{tx}, false
}

func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []Tx, token, adress string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, token, adress)
		if !ok {
			continue
		}
		txs = append(txs, tx...)
	}
	return
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.dexClient.GetAccountMetadata(address)
	if err != nil || len(account.Balances) == 0 {
		return []blockatlas.Token{}, nil
	}
	tokens, err := p.dexClient.GetTokens()
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Balances, tokens), nil
}

// NormalizeToken converts a Binance token into the generic model
func NormalizeToken(srcToken *Balance, tokens *TokenPage) (t blockatlas.Token, ok bool) {
	if srcToken.isAllZeroBalance() {
		return t, false
	}

	tk := tokens.findToken(srcToken.Symbol)
	if tk == nil {
		return t, false
	}

	t = blockatlas.Token{
		Name:     tk.Name,
		Symbol:   tk.OriginalSymbol,
		TokenID:  tk.Symbol,
		Coin:     coin.BNB,
		Decimals: uint(decimalPlaces(tk.TotalSupply)),
		Type:     blockatlas.TokenTypeBEP2,
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

package ethereum

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"math/big"
	"net/http"
	"strconv"
)

var (
	supportedTypes = map[string]bool{"ERC721": true, "ERC1155": true}
	slugTokens     = map[string]bool{"ERC1155": true}
)

type Platform struct {
	client            Client
	collectionsClient CollectionsClient
	CoinIndex         uint
}

func (p *Platform) Init() error {
	handle := coin.Coins[p.CoinIndex].Handle

	p.client.HTTPClient = http.DefaultClient
	p.client.BaseURL = viper.GetString(fmt.Sprintf("%s.api", handle))

	p.collectionsClient.HTTPClient = http.DefaultClient
	p.collectionsClient.CollectionsURL = viper.GetString(fmt.Sprintf("%s.collections_api", handle))
	p.collectionsClient.CollectionsApiKey = viper.GetString(fmt.Sprintf("%s.collections_api_key", handle))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
}

func (p *Platform) getTransactions(c *gin.Context) {
	token := c.Query("token")
	address := c.Param("address")
	var srcPage *Page
	var err error

	if token != "" {
		srcPage, err = p.client.GetTxsWithContract(address, token)
	} else {
		srcPage, err = p.client.GetTxs(address)
	}

	if apiError(c, err) {
		return
	}

	var txs []blockatlas.Tx
	for _, srcTx := range srcPage.Docs {
		txs = AppendTxs(txs, &srcTx, p.CoinIndex)
	}

	page := blockatlas.TxPage(txs)
	page.Sort()
	c.JSON(http.StatusOK, &page)
}

func extractBase(srcTx *Doc, coinIndex uint) (base blockatlas.Tx, ok bool) {
	var status blockatlas.Status
	var errReason string
	if srcTx.Error == "" {
		status = blockatlas.StatusCompleted
	} else {
		status = blockatlas.StatusFailed
		errReason = srcTx.Error
	}

	unix, err := strconv.ParseInt(srcTx.TimeStamp, 10, 64)
	if err != nil {
		return base, false
	}

	fee := calcFee(srcTx.GasPrice, srcTx.GasUsed)

	base = blockatlas.Tx{
		ID:       srcTx.ID,
		Coin:     coinIndex,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      blockatlas.Amount(fee),
		Date:     unix,
		Block:    srcTx.BlockNumber,
		Status:   status,
		Error:    errReason,
		Sequence: srcTx.Nonce,
	}
	return base, true
}

func AppendTxs(in []blockatlas.Tx, srcTx *Doc, coinIndex uint) (out []blockatlas.Tx) {
	out = in
	baseTx, ok := extractBase(srcTx, coinIndex)
	if !ok {
		return
	}

	// Native ETH transaction
	if len(srcTx.Ops) == 0 && srcTx.Input == "0x" {
		transferTx := baseTx
		transferTx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coinIndex].Symbol,
			Decimals: coin.Coins[coinIndex].Decimals,
		}
		out = append(out, transferTx)
	}

	// Smart Contract Call
	if len(srcTx.Ops) == 0 && srcTx.Input != "0x" {
		contractTx := baseTx
		contractTx.Meta = blockatlas.ContractCall{
			Input: srcTx.Input,
			Value: srcTx.Value,
		}
		out = append(out, contractTx)
	}

	if len(srcTx.Ops) == 0 {
		return
	}
	op := &srcTx.Ops[0]

	if op.Type == blockatlas.TxTokenTransfer && op.Contract != nil {
		tokenTx := baseTx

		tokenTx.Meta = blockatlas.TokenTransfer{
			Name:     op.Contract.Name,
			Symbol:   op.Contract.Symbol,
			TokenID:  op.Contract.Address,
			Decimals: op.Contract.Decimals,
			Value:    blockatlas.Amount(op.Value),
			From:     op.From,
			To:       op.To,
		}
		out = append(out, tokenTx)
	}
	return
}

func calcFee(gasPrice string, gasUsed string) string {
	var gasPriceBig, gasUsedBig, feeBig big.Int

	gasPriceBig.SetString(gasPrice, 10)
	gasUsedBig.SetString(gasUsed, 10)

	feeBig.Mul(&gasPriceBig, &gasUsedBig)

	return feeBig.String()
}

func apiError(c *gin.Context, err error) bool {
	if err != nil {
		logger.Error(err, "Unhandled error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return true
	}
	return false
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcPage, err := p.client.GetBlockByNumber(num); err == nil {
		var txs []blockatlas.Tx
		for _, srcTx := range srcPage {
			txs = AppendTxs(txs, &srcTx, p.CoinIndex)
		}
		return &blockatlas.Block{
			Number: num,
			ID:     strconv.FormatInt(num, 10),
			Txs:    txs,
		}, nil
	} else {
		return nil, err
	}
}

func (p *Platform) GetCollections(owner string) (blockatlas.CollectionPage, error) {
	collections, err := p.collectionsClient.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectionPage(collections, p.CoinIndex, owner)
	return page, nil
}

func (p *Platform) GetCollectibles(owner, collectibleID string) (blockatlas.CollectiblePage, error) {
	collection, items, err := p.collectionsClient.GetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectiblePage(collection, items, p.CoinIndex)
	return page, nil
}

func NormalizeCollectionPage(collections []Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
	for _, collection := range collections {
		if len(collection.Contracts) == 0 {
			continue
		}
		item := NormalizeCollection(collection, coinIndex, owner)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollection(c Collection, coinIndex uint, owner string) blockatlas.Collection {
	if len(c.Contracts) == 0 {
		return blockatlas.Collection{}
	}

	description := getValidParameter(c.Contracts[0].Description, c.Description)
	symbol := getValidParameter(c.Contracts[0].Symbol, "")
	collectionId := getValidParameter(c.Contracts[0].Address, "")
	version := getValidParameter(c.Contracts[0].NftVersion, "")
	collectionType := getValidParameter(c.Contracts[0].Type, "")
	if _, ok := slugTokens[collectionType]; ok {
		collectionId = createCollectionId(collectionId, c.Slug)
	}

	return blockatlas.Collection{
		Name:            c.Name,
		Symbol:          symbol,
		Slug:            c.Slug,
		ImageUrl:        c.ImageUrl,
		Description:     description,
		ExternalLink:    c.ExternalUrl,
		Total:           int(c.Total.Int64()),
		Id:              collectionId,
		CategoryAddress: collectionId,
		Address:         owner,
		Version:         version,
		Coin:            coinIndex,
		Type:            collectionType,
	}
}

func NormalizeCollectiblePage(c *Collection, srcPage []Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
	if len(c.Contracts) == 0 {
		return
	}
	for _, src := range srcPage {
		item := NormalizeCollectible(c, src, coinIndex)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollectible(c *Collection, a Collectible, coinIndex uint) blockatlas.Collectible {
	address := getValidParameter(c.Contracts[0].Address, "")
	collectionType := getValidParameter(c.Contracts[0].Type, "")
	collectionID := address
	if _, ok := slugTokens[collectionType]; ok {
		collectionID = createCollectionId(address, c.Slug)
	}
	externalLink := getValidParameter(a.ExternalLink, a.AssetContract.ExternalLink)
	return blockatlas.Collectible{
		CollectionID:     collectionID,
		ContractAddress:  address,
		TokenID:          a.TokenId,
		CategoryContract: a.AssetContract.Address,
		Name:             a.Name,
		Category:         c.Name,
		ImageUrl:         a.ImagePreviewUrl,
		ProviderLink:     a.Permalink,
		ExternalLink:     externalLink,
		Type:             collectionType,
		Description:      a.Description,
		Coin:             coinIndex,
	}
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, *p), nil
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Token, coinIndex uint) (t blockatlas.Token, ok bool) {
	t = blockatlas.Token{
		Name:     srcToken.Contract.Name,
		Symbol:   srcToken.Contract.Symbol,
		TokenID:  srcToken.Contract.Contract,
		Coin:     coinIndex,
		Decimals: srcToken.Contract.Decimals,
		Type:     blockatlas.TokenTypeERC20,
	}

	return t, true
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Token, p Platform) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0)
	for _, srcToken := range srcTokens {
		token, ok := NormalizeToken(&srcToken, p.CoinIndex)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}

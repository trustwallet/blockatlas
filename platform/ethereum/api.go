package ethereum

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	p.collectionsClient.CollectionsURL = viper.GetString(fmt.Sprintf("%s.opensea_api", handle))
	p.collectionsClient.CollectionsApiKey = viper.GetString(fmt.Sprintf("%s.opensea_api_key", handle))
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
	var status, errReason string
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
			Value: blockatlas.Amount(srcTx.Value),
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

	if op.Type == blockatlas.TxTokenTransfer {
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
		logrus.WithError(err).Errorf("Unhandled error")
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
	items, err := p.collectionsClient.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectionPage(items, p.CoinIndex, owner)
	return page, nil
}

func (p *Platform) GetCollectibles(owner string, collectibleID string) (blockatlas.CollectiblePage, error) {
	items, err := p.collectionsClient.GetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectiblePage(items, p.CoinIndex)
	return page, nil
}

func NormalizeCollectionPage(srcPage []Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
	for _, src := range srcPage {
		item := NormalizeCollection(src, coinIndex, owner)
		page = append(page, item)
	}
	return
}

func NormalizeCollection(c Collection, coinIndex uint, owner string) blockatlas.Collection {
	return blockatlas.Collection{
		Name:            c.Name,
		Symbol:          c.Contract[0].Symbol,
		ImageUrl:        c.ImageUrl,
		Description:     c.Contract[0].Description,
		ExternalLink:    c.ExternalUrl,
		Total:           c.Total,
		CategoryAddress: c.Contract[0].Address,
		Address:         owner,
		Version:         c.Contract[0].NftVersion,
		Coin:            coinIndex,
		Type:            c.Contract[0].Type,
	}
}

func NormalizeCollectiblePage(srcPage []Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
	for _, src := range srcPage {
		item := NormalizeCollectible(src, coinIndex)
		page = append(page, item)
	}
	return
}

func NormalizeCollectible(a Collectible, coinIndex uint) blockatlas.Collectible {
	return blockatlas.Collectible{
		TokenID:         a.TokenId,
		ContractAddress: a.AssetContract.Address,
		Name:            a.Name,
		Category:        a.AssetContract.Category,
		ImageUrl:        a.ImagePreviewUrl,
		ExternalLink:    a.ExternalLink,
		Type:            "ERC721",
		Description:     a.Description,
		Coin:            coinIndex,
	}
}

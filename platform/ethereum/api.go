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
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	handle := coin.Coins[p.CoinIndex].Handle
	p.client.BaseURL = viper.GetString(fmt.Sprintf("%s.api", handle))
	p.client.CollectionsURL = viper.GetString(fmt.Sprintf("%s.collections_api", handle))
	p.client.CollectionsApiKey = viper.GetString(fmt.Sprintf("%s.collections_api_key", handle))
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) RegisterRoutes(router gin.IRouter) {
	router.GET("/:address", func(c *gin.Context) {
		p.getTransactions(c)
	})
	router.GET("/:address/collections", func(c *gin.Context) {
		p.getCollections(c)
	})
	router.GET("/:address/collections/:contractAddress", func(c *gin.Context) {
		p.getCollectibles(c)
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

func (p *Platform) getCollections(c *gin.Context) {
	ownerAddress := c.Param("address")
	items, err := p.client.GetCollections(ownerAddress)
	if apiError(c, err) {
		return
	}
	page := NormalizeCollectionPage(items, p.CoinIndex)
	c.JSON(http.StatusOK, &page)
}

func (p *Platform) getCollectibles(c *gin.Context) {
	ownerAddress := c.Param("address")
	contractAddress := c.Param("contractAddress")
	items, err := p.client.GetCollectibles(ownerAddress, contractAddress)
	if apiError(c, err) {
		return
	}
	page := NormalizeCollectiblePage(items, p.CoinIndex)
	c.JSON(http.StatusOK, &page)
}

func NormalizeCollectionPage(srcPage []Collection, coinIndex uint) (page blockatlas.CollectionPage) {
	for _, src := range srcPage {
		item := NormalizeCollection(src, coinIndex)
		page = append(page, item)
	}
	return
}

func NormalizeCollection(c Collection, coinIndex uint) blockatlas.Collection {
	return blockatlas.Collection{
		Name:            c.Name,
		Symbol:          c.Contract[0].Symbol,
		ImageUrl:        c.ImageUrl,
		Description:     c.Contract[0].Description,
		ExternalLink:    c.ExternalUrl,
		Total:           strconv.Itoa(c.Total),
		CategoryAddress: c.Contract[0].Address,
		Address:         "",
		Version:         c.Contract[0].NftVersion,
		Coin:            coinIndex,
		Type:            "ERC721",
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
		ImageUrl:        a.ImageUrl,
		ExternalLink:    a.ExternalLink,
		Type:            "ERC721",
		Description:     a.Description,
		Coin:            coinIndex,
	}
}

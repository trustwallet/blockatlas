package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// @Summary Get Tokens
// @ID tokens
// @Description Get tokens from the address
// @Accept json
// @Produce json
// @Tags Transactions
// @Param coin path string true "the coin name" default(ethereum)
// @Param address path string true "the query address" default(0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB)
// @Success 200 {object} blockatlas.CollectionPage
// @Failure 500 {object} ErrorResponse
// @Router /v2/{coin}/tokens/{address} [get]
func GetTokensByAddress(c *gin.Context, tokenAPI blockatlas.TokensAPI) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusOK, blockatlas.TxPage{})
		return
	}

	result, err := tokenAPI.GetTokenListByAddress(address)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &result})
}

// @Description Get tokens
// @ID tokens_v3
// @Summary Get list of tokens by map: coin -> [addresses]
// @Accept json
// @Produce json
// @Tags Transactions
// @Param data body string true "Payload" default({"60": ["0xb3624367b1ab37daef42e1a3a2ced012359659b0"]})
// @Success 200 {object} blockatlas.ResultsResponse
// @Router /v2/tokens [post]
func GetTokens(c *gin.Context, apis map[uint]blockatlas.TokensAPI) {
	var query map[string][]string
	if err := c.Bind(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	result := make(blockatlas.TokenPage, 0)
	for coinStr, addresses := range query {
		coinNum, err := strconv.ParseUint(coinStr, 10, 32)
		if err != nil {
			continue
		}
		api, ok := apis[uint(coinNum)]
		if !ok {
			continue
		}

		tokens := getTokens(api, addresses)
		result = append(result, tokens...)
	}
	c.JSON(http.StatusOK, blockatlas.ResultsResponse{Total: len(result), Results: &result})
}

func getTokens(tokenAPI blockatlas.TokensAPI, addresses []string) blockatlas.TokenPage {
	var (
		tokenPagesChan = make(chan blockatlas.TokenPage, len(addresses))
		wg             sync.WaitGroup
		result         blockatlas.TokenPage
		timeout        = time.Second * 3
	)

	for _, address := range addresses {
		wg.Add(1)
		go func(address string, wg *sync.WaitGroup) {
			defer wg.Done()

			stopChan := make(chan struct{})
			pageChan := make(chan blockatlas.TokenPage)

			go func() {
				defer close(stopChan)
				defer close(pageChan)
				tokenPage, err := tokenAPI.GetTokenListByAddress(address)
				if err != nil {
					stopChan <- struct{}{}
					return
				}
				pageChan <- tokenPage
			}()

			select {
			case <-time.After(timeout):
				return
			case <-stopChan:
				return
			case p := <-pageChan:
				tokenPagesChan <- p
			}
		}(address, &wg)
	}
	wg.Wait()
	close(tokenPagesChan)

	for page := range tokenPagesChan {
		result = append(result, page...)
	}

	return result
}

func GetTokensByAddressIndexer(c *gin.Context, database *db.Instance) {
	var query map[string]string
	if err := c.Bind(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var addresses []string
	for coinID, a := range query {
		addresses = append(addresses, coinID+"_"+a)
	}
	assetsByAddresses, err := database.GetAssetsMapByAddresses(addresses, c.Request.Context())
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(errors.New("db issue")))
		return
	}
	c.JSON(http.StatusOK, assetsByAddresses)
}

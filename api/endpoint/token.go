package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/model"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/http"
	"strconv"
	"sync"
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
// @Failure 500 {object} middleware.ApiError
// @Router /v2/{coin}/tokens/{address} [get]
func GetTokensByAddress(c *gin.Context, tokenAPI blockatlas.TokenAPI) {
	address := c.Param("address")
	if address == "" {
		EmptyPage(c)
		return
	}

	result, err := tokenAPI.GetTokenListByAddress(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
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
// @Router /v3/tokens [post]
func GetTokens(c *gin.Context, tokensAPIs map[uint]blockatlas.TokenAPI) {
	var query map[string][]string
	if err := c.Bind(&query); err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	result := make(blockatlas.TokenPage, 0)
	for coinStr, addresses := range query {
		coinNum, err := strconv.ParseUint(coinStr, 10, 32)
		if err != nil {
			continue
		}
		tokenAPI, ok := tokensAPIs[uint(coinNum)]
		if !ok {
			continue
		}

		tokens := getTokens(tokenAPI, addresses)
		result = append(result, tokens...)
	}
	c.JSON(http.StatusOK, blockatlas.ResultsResponse{Total: len(result), Results: &result})
}

func getTokens(tokenAPI blockatlas.TokenAPI, addresses []string) blockatlas.TokenPage {
	var (
		tokenPagesChan = make(chan blockatlas.TokenPage, len(addresses))
		wg             sync.WaitGroup
		result         blockatlas.TokenPage
	)

	for _, address := range addresses {
		wg.Add(1)
		go func(address string, wg *sync.WaitGroup) {
			defer wg.Done()

			tokenPage, err := tokenAPI.GetTokenListByAddress(address)
			if err != nil {
				return
			}

			tokenPagesChan <- tokenPage
		}(address, &wg)
	}
	wg.Wait()
	close(tokenPagesChan)

	for page := range tokenPagesChan {
		result = append(result, page...)
	}

	return result
}

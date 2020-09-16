package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type (
	tokensResult struct {
		Result blockatlas.TokenPage
		mu     sync.Mutex
	}
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
	result := tokensResult{
		Result: make(blockatlas.TokenPage, 0),
		mu:     sync.Mutex{},
	}
	var wg sync.WaitGroup
	for coinStr, addresses := range query {
		coinNum, err := strconv.ParseUint(coinStr, 10, 32)
		if err != nil {
			continue
		}
		api, ok := apis[uint(coinNum)]
		if !ok {
			continue
		}
		wg.Add(1)
		go getTokens(api, addresses, &result, &wg)
	}
	wg.Wait()
	c.JSON(http.StatusOK, blockatlas.ResultsResponse{Total: len(result.Result), Results: &result.Result})
}

func getTokens(tokenAPI blockatlas.TokensAPI, addresses []string, data *tokensResult, wg *sync.WaitGroup) {
	var (
		tokenPagesChan = make(chan blockatlas.TokenPage, len(addresses))
		wgLocal        sync.WaitGroup
		timeout        = time.Second * 3
	)
	defer wg.Done()
	for _, address := range addresses {
		wgLocal.Add(1)
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
		}(address, &wgLocal)
	}
	wgLocal.Wait()
	close(tokenPagesChan)
	data.mu.Lock()
	for page := range tokenPagesChan {
		r := data.Result
		data.Result = append(r, page...)
	}
	data.mu.Unlock()
}

func GetTokensByAddressIndexer(c *gin.Context, instance tokensearcher.Instance) {
	var query tokensearcher.Request
	if err := c.Bind(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := instance.HandleTokensRequest(query, c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

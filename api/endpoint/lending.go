package endpoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func ServeProviders(c *gin.Context, apis map[string]blockatlas.LendingAPI) {
	p, err := getProviders(apis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &p})
}

func ServeRates(c *gin.Context, apis map[string]blockatlas.LendingAPI) {
	provider, ok := c.Params.Get("provider")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fatal: missing provider"})
		return
	}
	var req blockatlas.RatesRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	p, err := getRates(provider, req, apis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &p})
}

func ServeAccount(c *gin.Context, apis map[string]blockatlas.LendingAPI) {
	provider, ok := c.Params.Get("provider")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fatal: missing provider"})
		return
	}
	var req blockatlas.AccountRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	p, err := getAccounts(provider, req, apis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &p})
}

// GetProviders return provider info list
func getProviders(apis map[string]blockatlas.LendingAPI) (*[]blockatlas.LendingProvider, error) {
	ret := []blockatlas.LendingProvider{}
	for _, api := range apis {
		prov, err := api.GetProviderInfo()
		if err != nil {
			continue
		}
		ret = append(ret, prov)
	}
	return &ret, nil
}

// GetRates return rates info
func getRates(provider string, req blockatlas.RatesRequest, apis map[string]blockatlas.LendingAPI) (*blockatlas.RatesResponse, error) {
	api, ok := apis[provider]
	if !ok {
		return nil, fmt.Errorf("Unknown provider %v", provider)
	}
	rates, err := api.GetCurrentLendingRates(req.Assets)
	if err != nil {
		return nil, err
	}
	return &blockatlas.RatesResponse{Provider: provider, Rates: rates}, nil
}

// GetAccounts return account contract
func getAccounts(provider string, req blockatlas.AccountRequest, apis map[string]blockatlas.LendingAPI) (*[]blockatlas.AccountLendingContracts, error) {
	api, ok := apis[provider]
	if !ok {
		return nil, fmt.Errorf("Unknown provider %v", provider)
	}
	return api.GetAccountLendingContracts(req)
}

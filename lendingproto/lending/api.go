package lending

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/lendingproto/compound"
	"github.com/trustwallet/blockatlas/lendingproto/model"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Lending API
// As currently only Compuond is planned, API is not made entirely generic, but prepared for later generalization.

var compoundProvider *compound.Provider

// Init Setup HTTP API
func Init(endpoint string) error {
	compoundProvider = compound.Init("https://api.compound.finance/api")

	r := gin.Default()

	r.GET("/v1/lending/providers", serveProviders)
	r.POST("/v1/lending/rates/:provider", serveRates)
	r.POST("/v1/lending/account/:provider", serveAccount)

	return r.Run(endpoint)
}

func serveProviders(c *gin.Context) {
	p, err := GetProviders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &p})
}

func serveRates(c *gin.Context) {
	provider, ok := c.Params.Get("provider")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fatal: missing provider"})
		return
	}
	var req model.RatesRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	p, err := GetRates(provider, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &p})
}

func serveAccount(c *gin.Context) {
	provider, ok := c.Params.Get("provider")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fatal: missing provider"})
		return
	}
	var req model.AccountRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	p, err := GetAccounts(provider, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockatlas.DocsResponse{Docs: &p})
}

// GetProviders return provider info list
func GetProviders() (*[]model.LendingProvider, error) {
	// we have one provider
	provCompound, err := compoundProvider.GetProviderInfo()
	if err != nil {
		return nil, err
	}
	return &[]model.LendingProvider{provCompound}, nil
}

// GetRates return rates info
func GetRates(provider string, req model.RatesRequest) (*model.RatesResponse, error) {
	// we have one provider
	if provider != compoundProvider.Name() {
		return nil, fmt.Errorf("Unknown provider %v", provider)
	}
	rates, err := compoundProvider.GetCurrentLendingRates(req.Assets)
	if err != nil {
		return nil, err
	}
	return &model.RatesResponse{Provider: provider, Rates: rates}, nil
}

// GetAccounts return account contract
func GetAccounts(provider string, req model.AccountRequest) (*[]model.AccountLendingContracts, error) {
	// we have one provider
	if provider != compoundProvider.Name() {
		return nil, fmt.Errorf("Unknown provider %v", provider)
	}
	return compoundProvider.GetAccountLendingContracts(req)
}

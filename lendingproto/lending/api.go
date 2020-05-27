package lending

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/lendingproto/compound"
	"github.com/trustwallet/blockatlas/lendingproto/model"
)

// Lending API
// As currently only Compuond is planned, API is not made entirely generic, but prepared for later generalization.

// Init Setup HTTP API
func Init(endpoint string) {
	r := gin.Default()

	r.GET("/v1/lending/providers", serveProviders)
	r.POST("/v1/lending/rates", serveRates)

	r.Run(endpoint)
}

func serveProviders(c *gin.Context) {
	p, err := GetProviders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, p)
}

func serveRates(c *gin.Context) {
	bodyB, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Fatal: " + err.Error()})
		return
	}
	var req model.RatesRequest
	err = json.Unmarshal(bodyB, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Parsing: " + err.Error()})
		return
	}
	p, err := GetRates(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, p)
}

var compoundProviderName string = "compound"

// GetProviders return provider info list
func GetProviders() (*[]model.LendingProvider, error) {
	// we have one provider
	provCompound, err := compound.GetProviderInfo()
	if err != nil {
		return nil, err
	}
	return &[]model.LendingProvider{provCompound}, nil
}

// GetRates return rates info
func GetRates(req model.RatesRequest) (*model.RatesResponse, error) {
	// we have one provider
	if req.Provider != compoundProviderName {
		return nil, fmt.Errorf("Unknown provider %v", req.Provider)
	}
	rates, err := compound.GetCurrentLendingRates(req.Assets)
	if err != nil {
		return nil, err
	}
	return &model.RatesResponse{req.Provider, rates}, nil
}

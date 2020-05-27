package lending

import (
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
	r.POST("/v1/lending/rates", serveProviders)

	r.Run(endpoint)
}

func serveProviders(c *gin.Context) {
	p, err := GetProviders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// GetProviders return provider info list
func GetProviders() ([]model.LendingProvider, error) {
	// we have one provider
	provCompound, err := compound.GetProviderInfo()
	if err != nil {
		return []model.LendingProvider{}, err
	}
	return []model.LendingProvider{provCompound}, nil
}

package main

import (
	"github.com/gin-gonic/gin"
)

// Lending API
// As currently only Compuond is planned, API is not made entirely generic, but prepared for later generalization.

func Init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/v1/lending/providers", ServeProviders)
	r.Run(":8080")
}

func ServeProviders(c *gin.Context) {
	p, err := GetProviders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

func GetProviders() ([]LendingProvider, error) {
	// we have one provider
	provCompound, err := GetProviderInfo()
	if err != nil {
		return []LendingProvider{}, err
	}
	return []LendingProvider{provCompound}, nil
}

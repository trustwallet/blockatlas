package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func getRoot(c *gin.Context) {
	c.String(http.StatusOK,
`Welcome to the Block Atlas API!

Don't know how you landed here?
Visit https://trustwallet.com to get back to the main page.

If you know what you're doing:
 - Visit /v1/ to list platforms
 - Source: https://github.com/trustwallet/blockatlas
 - Any questions? https://t.me/walletcore
`)
}

func getEnabledEndpoints(c *gin.Context) {
	var resp struct {
		Endpoints []string `json:"endpoints"`
	}
	for ns := range loaders {
		key := ns + ".api"
		if !viper.IsSet(key) || viper.GetString(key) == "" {
			continue
		}
		resp.Endpoints = append(resp.Endpoints, ns)
	}
	c.JSON(http.StatusOK, &resp)
}

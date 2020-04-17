package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"net/http"
)

func Init(serviceRepo *servicerepo.ServiceRepo) {
	initAssetsService(serviceRepo)
	initDomainsService(serviceRepo)
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"build":  internal.Build,
		"date":   internal.Date,
	})
}

func EmptyPage(c *gin.Context) {
	var page blockatlas.TxPage
	c.JSON(http.StatusOK, &page)
}

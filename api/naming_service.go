package api

import (
	"net/http"
	"strconv"
	"strings"

	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

type ZNSResponse struct {
	Addresses map[string]string
	Meta      struct {
		Owner string `json:"owner"`
		Type  string `json:"type"`
	} `json:"meta"`
}

// @Summary Lookup .eth / .zil addresses
// @ID lookup
// @Description Lookup ENS/ZNS to find registered addresses
// @Produce json
// @Tags ns
// @Param name query string empty "string name"
// @Param coin query string 60 "string coin"
// @Success 200 {object} blockatlas.Resolved
// @Failure 500 {object} api.ApiError
// @Router /ns/lookup [get]
func MakeLookupRoute(router gin.IRouter) {
	ns := router.Group("/ns")
	ns.GET("/lookup", handleLookup)
}

func handleLookup(c *gin.Context) {
	name := c.Query("name")
	coinQuery := c.DefaultQuery("coin", strconv.Itoa(CoinType.ETH))

	if name == "" {
		RenderError(c, http.StatusBadRequest, "name query is missing")
		return
	}
	coin, err := strconv.ParseUint(coinQuery, 10, 64)
	if err != nil {
		RenderError(c, http.StatusBadRequest, "coin query is invalid")
		return
	}
	if strings.Contains(name, ".eth") {
		handleENSLookup(c, name, coin)
	} else if strings.Contains(name, ".zil") {
		handleZILLookup(c, name, coin)
	} else {
		RenderError(c, http.StatusBadRequest, "not supported domain")
	}
}

func handleENSLookup(c *gin.Context, name string, coin uint64) {
	client, err := ethclient.Dial(viper.GetString("ethereum.rpc"))
	if err != nil {
		RenderError(c, http.StatusInternalServerError, "can't dial to ethereum rpc")
		return
	}
	defer client.Close()
	result := blockatlas.Resolved{
		Coin: coin,
	}

	ensName, err := ens.NewName(client, name)
	if err != nil {
		RenderError(c, http.StatusInternalServerError, err.Error())
		return
	}

	address, err := ensName.Address(coin)
	if err != nil {
		RenderError(c, http.StatusInternalServerError, err.Error())
		return
	}

	result.Result = mapAddress(coin, address)
	RenderSuccess(c, &result)
}

func mapAddress(coin uint64, bytes []byte) string {
	// FIXME: convert bytes to string according to coin
	address := hex.EncodeToString(bytes)
	if address != "" {
		address = "0x" + address
	}
	return address
}

func handleZILLookup(c *gin.Context, name string, coin uint64) {
	client := blockatlas.InitClient(viper.GetString("zilliqa.lookup"))
	var resp ZNSResponse
	client.Get(&resp, "/"+name, nil)
	result := blockatlas.Resolved{
		Coin: coin,
	}
	symbol := CoinType.Coins[uint(coin)].Symbol
	result.Result = resp.Addresses[symbol]
	RenderSuccess(c, &result)
}

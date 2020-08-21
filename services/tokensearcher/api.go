package tokensearcher

import (
	"context"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/watchmarket/pkg/watchmarket"
	"strconv"
	"strings"
	"sync"
)

type Instance struct {
	database *db.Instance
	apis     map[uint]blockatlas.TokensAPI
}

func Init(database *db.Instance, apis map[uint]blockatlas.TokensAPI) Instance {
	return Instance{database: database, apis: apis}
}

func (i Instance) HandleTokensRequest(request map[string][]string, ctx context.Context) (map[string][]string, error) {
	addresses := getAddressesFromRequest(request)
	assetsByAddresses, err := i.database.GetAssetsMapByAddresses(addresses, ctx)
	if err != nil {
		return nil, err
	}
	addressesToRegisterByCoin := getAddressesToRegisterByCoin(assetsByAddresses, addresses)
	assetsByAddressesToRegister := getAssetsForAddressesFromNodes(addressesToRegisterByCoin, i.apis)
	err = publishNewAddressesToQueue(assetsByAddressesToRegister)
	if err != nil {
		logger.Error(err)
	}
	return getAssetsToResponse(assetsByAddresses, assetsByAddressesToRegister, addresses), nil
}

func getAddressesFromRequest(request map[string][]string) []string {
	var addresses []string
	for coinID, requestAddresses := range request {
		for _, a := range requestAddresses {
			addresses = append(addresses, coinID+"_"+a)
		}
	}
	return addresses
}

func getAddressesToRegisterByCoin(assetsByAddresses map[string][]string, addressesFromRequest []string) map[uint][]string {
	addressesToRegisterByCoin := make(map[uint][]string)
	addressesFromRequestMap := make(map[string]bool)
	for _, a := range addressesFromRequest {
		addressesFromRequestMap[a] = true
	}
	for _, address := range addressesFromRequest {
		_, ok := assetsByAddresses[address]
		if !ok {
			a, coinID, ok := getCoinIDFromAddress(address)
			if !ok {
				continue
			}
			currentAddresses := addressesToRegisterByCoin[coinID]
			addressesToRegisterByCoin[coinID] = append(currentAddresses, a)
		}
	}
	return addressesToRegisterByCoin
}

func getAssetsForAddressesFromNodes(addresses map[uint][]string, apis map[uint]blockatlas.TokensAPI) map[string][]string {
	a := assetsByAddresses{Result: make(map[string][]string)}
	var wg sync.WaitGroup
	for coinID, addresses := range addresses {
		api, ok := apis[coinID]
		if !ok {
			continue
		}
		wg.Add(1)
		go fetchAssetsByAddresses(api, addresses, &a, &wg)
	}
	wg.Wait()
	return a.Result
}

func publishNewAddressesToQueue(map[string][]string) error {
	return nil
}

func getCoinIDFromAddress(address string) (string, uint, bool) {
	result := strings.Split(address, "_")
	if len(result) != 2 {
		return "", 0, false
	}
	id, err := strconv.Atoi(result[0])
	if err != nil {
		return "", 0, false
	}
	return result[1], uint(id), true
}

func fetchAssetsByAddresses(tokenAPI blockatlas.TokensAPI, addresses []string, result *assetsByAddresses, wg *sync.WaitGroup) {
	for _, a := range addresses {
		tokens, err := tokenAPI.GetTokenListByAddress(a)
		if err != nil {
			logger.Error("Chain: " + tokenAPI.Coin().Handle + " Address: " + a)
			continue
		}
		result.Lock()
		for _, t := range tokens {
			r := result.Result[a]
			result.Result[a] = append(r, watchmarket.BuildID(t.Coin, t.TokenID))
		}
		result.Unlock()
	}
	wg.Done()
}

func getAssetsToResponse(assetsFromDB, assetsFromNodes map[string][]string, addressesFromRequest []string) map[string][]string {
	result := make(map[string][]string)
	for _, address := range addressesFromRequest {
		assetsFromDBForAddress, ok := assetsFromDB[address]
		if !ok {
			assetsFromNodesForAddress, ok := assetsFromNodes[address]
			if !ok {
				continue
			}
			result[address] = assetsFromNodesForAddress
			continue
		}
		result[address] = assetsFromDBForAddress
	}
	return result
}

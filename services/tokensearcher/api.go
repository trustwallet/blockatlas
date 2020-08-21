package tokensearcher

import (
	"context"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
	var addresses []string
	for coinID, requestAddresses := range request {
		for _, a := range requestAddresses {
			addresses = append(addresses, coinID+"_"+a)
		}
	}
	assetsByAddressesFromDB, err := i.database.GetAssetsMapByAddresses(addresses, ctx)
	if err != nil {
		return nil, err
	}

	addressesToRegisterByCoin := getAddressesToRegisterByCoin(assetsByAddressesFromDB, addresses)
	assetsByAddressesToRegister := getAssetsForAddressesFromNodes(addressesToRegisterByCoin, i.apis)

	err = publishNewAddressesToQueue(assetsByAddressesToRegister)
	if err != nil {
		return nil, err
	}

	return getAssetsToResponse(assetsByAddressesFromDB, assetsByAddressesToRegister), nil
}

func getAddressesToRegisterByCoin(assetsByAddresses map[string][]string, addressesFromRequest []string) map[uint][]string {
	addressesToRegisterByCoin := make(map[uint][]string)
	addressesFromRequestMap := make(map[string]bool)
	for _, a := range addressesFromRequest {
		addressesFromRequestMap[a] = true
	}
	for address, _ := range assetsByAddresses {
		_, ok := addressesFromRequestMap[address]
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

}

func getAssetsToResponse(assetsFromDB, assetsFromNodes map[string][]string) map[string][]string {
	return nil
}

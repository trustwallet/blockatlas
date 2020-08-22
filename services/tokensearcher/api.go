package tokensearcher

import (
	"context"
	"encoding/json"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strconv"
	"strings"
	"sync"
)

type (
	AddressesByCoin map[uint][]string
	AssetsByAddress map[string][]string
	Request         map[string][]string
)

type Instance struct {
	database *db.Instance
	apis     map[uint]blockatlas.TokensAPI
	queue    mq.Queue
}

func Init(database *db.Instance, apis map[uint]blockatlas.TokensAPI, queue mq.Queue) Instance {
	return Instance{database: database, apis: apis, queue: queue}
}

func (i Instance) HandleTokensRequest(request Request, ctx context.Context) (AssetsByAddress, error) {
	addresses := getAddressesFromRequest(request)
	dbAssetsMap, err := i.database.GetAssetsMapByAddresses(addresses, ctx)
	if err != nil {
		return nil, err
	}
	assetsByAddress := make(AssetsByAddress)
	addressesByCoin := getAddressesToRegisterByCoin(dbAssetsMap, addresses)
	if len(addressesByCoin) == 0 {
		assetsByAddress = getAssetsByAddressFromNodes(addressesByCoin, i.apis)
		err = publishNewAddressesToQueue(i.queue, assetsByAddress)
		if err != nil {
			logger.Error(err)
		}
	}

	return getAssetsToResponse(dbAssetsMap, assetsByAddress, addresses), nil
}

func getAddressesFromRequest(request Request) []string {
	var addresses []string
	for coinID, requestAddresses := range request {
		for _, a := range requestAddresses {
			addresses = append(addresses, coinID+"_"+a)
		}
	}
	return addresses
}

func getAddressesToRegisterByCoin(assetsByAddresses AssetsByAddress, addresses []string) AddressesByCoin {
	addressesByCoin := make(AddressesByCoin)
	addressesFromRequestMap := make(map[string]bool)
	for _, a := range addresses {
		addressesFromRequestMap[a] = true
	}
	for _, address := range addresses {
		_, ok := assetsByAddresses[address]
		if !ok {
			a, coinID, ok := getCoinIDFromAddress(address)
			if !ok {
				continue
			}
			currentAddresses := addressesByCoin[coinID]
			addressesByCoin[coinID] = append(currentAddresses, a)
		}
	}
	return addressesByCoin
}

func getAssetsByAddressFromNodes(addressesByCoin AddressesByCoin, apis map[uint]blockatlas.TokensAPI) AssetsByAddress {
	a := NodesResponse{AssetsByAddress: make(AssetsByAddress)}
	var wg sync.WaitGroup
	for coinID, addresses := range addressesByCoin {
		api, ok := apis[coinID]
		if !ok {
			continue
		}
		wg.Add(1)
		go fetchAssetsByAddresses(api, addresses, &a, &wg)
	}
	wg.Wait()
	return a.AssetsByAddress
}

func fetchAssetsByAddresses(tokenAPI blockatlas.TokensAPI, addresses []string, result *NodesResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	var tWg sync.WaitGroup
	tWg.Add(len(addresses))
	for _, a := range addresses {
		go func(address string, tWg *sync.WaitGroup) {
			defer tWg.Done()
			tokens, err := tokenAPI.GetTokenListByAddress(address)
			if err != nil {
				logger.Error("Chain: " + tokenAPI.Coin().Handle + " Address: " + address)
				return
			}
			result.UpdateAssetsByAddress(tokens, int(tokenAPI.Coin().ID), address)
		}(a, &tWg)
	}
	tWg.Wait()
}

func publishNewAddressesToQueue(queue mq.Queue, message AssetsByAddress) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return queue.Publish(body)
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

func getAssetsToResponse(dbAssetsMap, nodesAssetsMap AssetsByAddress, addresses []string) map[string][]string {
	result := make(map[string][]string)
	for _, address := range addresses {
		dbAddresses, ok := dbAssetsMap[address]
		if !ok {
			nodesAssets, ok := nodesAssetsMap[address]
			if !ok {
				continue
			}
			result[address] = nodesAssets
			continue
		}
		result[address] = dbAddresses
	}
	return result
}

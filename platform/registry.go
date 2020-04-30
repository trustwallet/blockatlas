package platform

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

var (
	// Platforms contains all registered platforms by handle
	Platforms map[string]blockatlas.Platform

	// BlockAPIs contain platforms with block services
	BlockAPIs map[string]blockatlas.BlockAPI

	// TxByAddrAPIs contains handlers with address-based transactions service
	TxByAddrAPIs map[string]blockatlas.TxAPI

	// TxByAddrAndXPubAPIs contains handlers with Xaddress- and PUB-based transactions service
	TxByAddrAndXPubAPIs map[string]blockatlas.TxByAddrAndXPubAPI

	// TokensAPIs contain platforms with token services
	TokensAPIs map[uint]blockatlas.TokensAPI

	// StakeAPIs contain platforms with staking services
	StakeAPIs map[string]blockatlas.StakeAPI

	// CollectionsAPIs contain platforms which collections services
	CollectionsAPIs map[uint]blockatlas.CollectionsAPI

	// NamingAPIs contain platforms which support naming services
	NamingAPIs map[uint64]blockatlas.NamingServiceAPI
)

func getActivePlatforms(handle string) []blockatlas.Platform {
	platforms := getPlatformMap()
	logger.Info("Platform API setup with: ", logger.Params{"handle": handle})

	if handle == allPlatformsHandle {
		return platforms.GetPlatformList()
	}

	platform, ok := platforms[handle]
	if ok {
		return []blockatlas.Platform{platform}
	}

	logger.Fatal("Please, use ATLAS_PLATFORM handle with non-empty value, see more at Readme. Example: all", logger.Params{"ATLAS_PLATFORM": handle})
	return nil
}

func Init(platformHandle string) {
	platformList := getActivePlatforms(platformHandle)

	// white list of collection api coins (only ETH now)
	InitCollectionsWhitelist()

	Platforms = make(map[string]blockatlas.Platform)
	BlockAPIs = make(map[string]blockatlas.BlockAPI)
	TxByAddrAPIs        = make(map[string]blockatlas.TxAPI)
	TxByAddrAndXPubAPIs = make(map[string]blockatlas.TxByAddrAndXPubAPI)
	TokensAPIs = make(map[uint]blockatlas.TokensAPI)
	StakeAPIs = make(map[string]blockatlas.StakeAPI)
	CollectionsAPIs = make(map[uint]blockatlas.CollectionsAPI)
	NamingAPIs = make(map[uint64]blockatlas.NamingServiceAPI)

	for _, platform := range platformList {
		handle := platform.Coin().Handle
		apiURL := fmt.Sprintf("%s.api", handle)

		if !viper.IsSet(apiURL) {
			continue
		}
		if viper.GetString(apiURL) == "" {
			continue
		}

		p := logger.Params{
			"platform": handle,
			"coin":     platform.Coin(),
		}

		if _, exists := Platforms[handle]; exists {
			logger.Fatal("Duplicate handle", p)
		}
		Platforms[handle] = platform
		if blockAPI, ok := platform.(blockatlas.BlockAPI); ok {
			BlockAPIs[handle] = blockAPI
		}
		if txByAddrAndXPubAPI, ok := platform.(blockatlas.TxByAddrAndXPubAPI); ok {
			TxByAddrAndXPubAPIs[handle] = txByAddrAndXPubAPI
		} else {
			// ByAddrAndXPub prevents ByAddr API
			if txByAddrAPI, ok := platform.(blockatlas.TxAPI); ok {
				TxByAddrAPIs[handle] = txByAddrAPI
			}
		}
		if tokenAPI, ok := platform.(blockatlas.TokensAPI); ok {
			TokensAPIs[platform.Coin().ID] = tokenAPI
		}
		if stakeAPI, ok := platform.(blockatlas.StakeAPI); ok {
			StakeAPIs[handle] = stakeAPI
		}
		if collectionAPI, ok := platform.(blockatlas.CollectionsAPI); ok && CollectionsWhitelist[platform.Coin().ID] {
			CollectionsAPIs[platform.Coin().ID] = collectionAPI
		}
		if namingAPI, ok := platform.(blockatlas.NamingServiceAPI); ok {
			NamingAPIs[uint64(platform.Coin().ID)] = namingAPI
		}
	}
}

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

	// StakeAPIs contain platforms with staking services
	StakeAPIs map[string]blockatlas.StakeAPI

	// CustomAPIs contain platforms with custom HTTP services
	CustomAPIs map[string]blockatlas.CustomAPI

	// NamingAPIs contain platforms which support naming services
	NamingAPIs map[uint64]blockatlas.NamingServiceAPI

	// CollectionAPIs contain platforms which collections services
	CollectionAPIs map[uint]blockatlas.CollectionAPI
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
	StakeAPIs = make(map[string]blockatlas.StakeAPI)
	CustomAPIs = make(map[string]blockatlas.CustomAPI)
	NamingAPIs = make(map[uint64]blockatlas.NamingServiceAPI)
	CollectionAPIs = make(map[uint]blockatlas.CollectionAPI)

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
		if stakeAPI, ok := platform.(blockatlas.StakeAPI); ok {
			StakeAPIs[handle] = stakeAPI
		}
		if customAPI, ok := platform.(blockatlas.CustomAPI); ok {
			CustomAPIs[handle] = customAPI
		}
		if namingAPI, ok := platform.(blockatlas.NamingServiceAPI); ok {
			NamingAPIs[uint64(platform.Coin().ID)] = namingAPI
		}
		if collectionAPI, ok := platform.(blockatlas.CollectionAPI); ok && CollectionsWhitelist[platform.Coin().ID] {
			CollectionAPIs[platform.Coin().ID] = collectionAPI
		}
	}
}

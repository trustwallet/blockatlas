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

	// TokensAPIs contain platforms with token services
	TokensAPIs map[uint]blockatlas.TokensAPI

	// StakeAPIs contain platforms with staking services
	StakeAPIs map[string]blockatlas.StakeAPI

	// CollectionsAPIs contain platforms which collections services
	CollectionsAPIs blockatlas.CollectionsAPIs

	// NamingAPIs contain platforms which support naming services
	NamingAPIs map[uint]blockatlas.NamingServiceAPI
)

func getActivePlatforms(handle string) []blockatlas.Platform {
	platforms := getAllHandlers()
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

	Platforms = make(map[string]blockatlas.Platform)
	BlockAPIs = make(map[string]blockatlas.BlockAPI)
	TokensAPIs = make(map[uint]blockatlas.TokensAPI)
	StakeAPIs = make(map[string]blockatlas.StakeAPI)

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
		if tokenAPI, ok := platform.(blockatlas.TokensAPI); ok {
			TokensAPIs[platform.Coin().ID] = tokenAPI
		}
		if stakeAPI, ok := platform.(blockatlas.StakeAPI); ok {
			StakeAPIs[handle] = stakeAPI
		}
	}

	CollectionsAPIs = getCollectionsHandlers()
	NamingAPIs = getNamingHandlers()
}

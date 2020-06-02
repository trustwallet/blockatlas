package platform

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/ethereum/compound"
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

	// LendingAPI contains lending providers, key is provider name
	LendingAPIs map[string]blockatlas.LendingAPI
)

func getActivePlatforms(handles []string) []blockatlas.Platform {
	if len(handles) == 0 {
		logger.Fatal("Please, use ATLAS_PLATFORM handle with non-empty value, see more at Readme. Example: all", logger.Params{"ATLAS_PLATFORM": handles})
		return nil
	}

	allPlatforms := getAllHandlers()
	logger.Info("Platform API setup with: ", logger.Params{"handles": handles})

	platforms := make([]blockatlas.Platform, 0, len(handles))

	for _, handle := range handles {
		if handle == allPlatformsHandle {
			return allPlatforms.GetPlatformList()
		}
		p, ok := allPlatforms[handle]
		if ok {
			platforms = append(platforms, p)
		}
	}
	return platforms
}

func Init(platformHandles []string) {
	platformList := getActivePlatforms(platformHandles)

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

	compoundLendingProvider := compound.Init("https://api.compound.finance/api") // TODO into config
	LendingAPIs = make(map[string]blockatlas.LendingAPI, 10)
	LendingAPIs[compoundLendingProvider.Name()] = compoundLendingProvider
}

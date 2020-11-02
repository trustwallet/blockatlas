package platform

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
)

func getActivePlatforms(handles []string) []blockatlas.Platform {
	if len(handles) == 0 {
		log.WithFields(log.Fields{"ATLAS_PLATFORM": handles}).
			Fatal("Please, use ATLAS_PLATFORM handle with non-empty value, see more at Readme. Example: all")
		return nil
	}

	allPlatforms := getAllHandlers()
	log.WithFields(log.Fields{"handles": handles}).Info("Platform API setup with")

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

		p := log.Fields{
			"platform": handle,
			"coin":     platform.Coin(),
		}

		if _, exists := Platforms[handle]; exists {
			log.WithFields(p).Fatal("Duplicate handle")
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
}

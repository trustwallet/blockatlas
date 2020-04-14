package blockatlas

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
)

type (
	// Platform can be used to access a crypto service
	Platform interface {
		Coin() coin.Coin
	}

	// TxAPI provides transaction lookups
	TxAPI interface {
		Platform
		GetTxsByAddress(address string) (TxPage, error)
	}

	// TokenTxAPI provides token transaction lookups
	TokenTxAPI interface {
		Platform
		GetTokenTxsByAddress(address, token string) (TxPage, error)
	}

	// TokenAPI provides token lookups
	TokenAPI interface {
		Platform
		GetTokenListByAddress(address string) (TokenPage, error)
	}

	// BlockAPI provides block information and lookups
	BlockAPI interface {
		Platform
		CurrentBlockNumber() (int64, error)
		GetBlockByNumber(num int64) (*Block, error)
	}

	// AddressAPI provides address information
	AddressAPI interface {
		Platform
		GetAddressesFromXpub(xpub string) ([]string, error)
	}

	// StakingAPI provides staking information
	StakeAPI interface {
		Platform
		UndelegatedBalance(address string) (string, error)
		GetDetails() StakingDetails
		GetValidators() (ValidatorPage, error)
		GetDelegations(address string) (DelegationsPage, error)
	}

	CollectionAPI interface {
		Platform
		GetCollections(owner string) (CollectionPage, error)
		GetCollectibles(owner, collectibleID string) (CollectiblePage, error)

		GetCollectionsV3(owner string) (CollectionPageV3, error)
		GetCollectiblesV3(owner, collectibleID string) (CollectiblePageV3, error)
	}

	// CustomAPI provides custom HTTP routes
	CustomAPI interface {
		Platform
		RegisterRoutes(router gin.IRouter)
	}

	// NamingServiceAPI provides public name service domains HTTP routes
	NamingServiceAPI interface {
		Platform
		Lookup(coins []uint64, name string) ([]Resolved, error)
	}
)

type Platforms map[string]Platform

func (ps Platforms) GetPlatformList() []Platform {
	platforms := make([]Platform, 0)
	for _, p := range ps {
		platforms = append(platforms, p)
	}
	return platforms
}

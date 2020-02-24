package blockatlas

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
)

// Platform can be used to access a crypto service
type Platform interface {
	Coin() coin.Coin
}

// TxAPI provides transaction lookups
type TxAPI interface {
	Platform
	GetTxsByAddress(address string) (TxPage, error)
}

// TokenTxAPI provides token transaction lookups
type TokenTxAPI interface {
	Platform
	GetTokenTxsByAddress(address, token string) (TxPage, error)
}

// TokenAPI provides token lookups
type TokenAPI interface {
	Platform
	GetTokenListByAddress(address string) (TokenPage, error)
}

// BlockAPI provides block information and lookups
type BlockAPI interface {
	Platform
	CurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64) (*Block, error)
}

// AddressAPI provides address information
type AddressAPI interface {
	Platform
	GetAddressesFromXpub(xpub string) ([]string, error)
}

// StakingAPI provides staking information
type StakeAPI interface {
	Platform
	UndelegatedBalance(address string) (string, error)
	GetDetails() StakingDetails
	GetValidators() (ValidatorPage, error)
	GetDelegations(address string) (DelegationsPage, error)
}

type CollectionAPI interface {
	Platform
	GetCollections(owner string) (CollectionPage, error)
	GetCollectibles(owner, collectibleID string) (CollectiblePage, error)

	//TODO: remove once most of the clients will be updated (deadline: March 17th)
	OldGetCollections(owner string) (CollectionPage, error)
	OldGetCollectibles(owner, collectibleID string) (CollectiblePage, error)

	GetCollectionsV4(owner string) (CollectionPage, error)
	GetCollectiblesV4(owner, collectibleID string) (CollectiblePage, error)
}

// CustomAPI provides custom HTTP routes
type CustomAPI interface {
	Platform
	RegisterRoutes(router gin.IRouter)
}

// NamingServiceAPI provides public name service domains HTTP routes
type NamingServiceAPI interface {
	Platform
	Lookup(coins []uint64, name string) ([]Resolved, error)
}

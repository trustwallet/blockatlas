package blockatlas

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

type (
	// Platform can be used to access a crypto service
	Platform interface {
		Coin() coin.Coin
	}

	// BlockAPI provides block information and lookups
	BlockAPI interface {
		Platform
		CurrentBlockNumber() (int64, error)
		GetBlockByNumber(num int64) (*types.Block, error)
	}

	// TxAPI provides transaction lookups based on address
	TxAPI interface {
		Platform
		GetTxsByAddress(address string) (types.Txs, error)
	}

	// TokenTxAPI provides token transaction lookups
	TokenTxAPI interface {
		Platform
		GetTokenTxsByAddress(address, token string) (types.Txs, error)
	}

	// TxUtxoAPI provides transaction lookup based on address and XPUB (Bitcoin-style)
	TxUtxoAPI interface {
		TxAPI
		GetTxsByXpub(xpub string) (types.Txs, error)
	}

	// TokensAPI provides token lookups
	TokensAPI interface {
		Platform
		GetTokenListByAddress(address string) ([]types.Token, error)
		GetTokenListIdsByAddress(address string) ([]string, error)
	}

	// StakingAPI provides staking information
	StakeAPI interface {
		Platform
		UndelegatedBalance(address string) (string, error)
		GetDetails() StakingDetails
		GetValidators() (ValidatorPage, error)
		GetDelegations(address string) (DelegationsPage, error)
		GetActiveValidators() (StakeValidators, error)
	}

	CollectionsAPI interface {
		Platform
		GetCollections(owner string) (types.CollectionPage, error)
		GetCollectibles(owner, collectibleID string) (types.CollectiblePage, error)
	}

	Platforms map[string]Platform

	CollectionsAPIs map[uint]CollectionsAPI
)

func (ps Platforms) GetPlatformList() []Platform {
	platforms := make([]Platform, 0)
	for _, p := range ps {
		platforms = append(platforms, p)
	}
	return platforms
}

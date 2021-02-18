package ethereum

import "github.com/trustwallet/golibs/types"

type EthereumClient interface {
	GetTransactions(address string, coinIndex uint) (types.Txs, error)
	GetTokenTxs(address, token string, coinIndex uint) (types.Txs, error)
	GetTokenList(address string, coinIndex uint) ([]types.Token, error)
	GetCurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64, coinIndex uint) (*types.Block, error)
}

type CollectibleClient interface {
	GetCollections(owner string, coinIndex uint) (types.CollectionPage, error)
	GetCollectibles(owner, collectionID string, coinIndex uint) (types.CollectiblePage, error)
}

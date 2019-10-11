package storage

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64) error
}

type Addresses interface {
	Lookup(coin uint, addresses ...string) ([]Subscription, error)
	AddSubscriptions([]interface{}) error
	DeleteSubscriptions([]interface{}) error
	GetAddressFromXpub(coin uint, xpub string) ([]Xpub, error)
	GetXpubFromAddress(coin uint, address string) (string, error)
	SaveXpubAddresses(coin uint, addresses []string, xpub string) error
}

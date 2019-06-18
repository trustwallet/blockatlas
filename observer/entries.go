package observer

type Subscription struct{
	Coin    uint    `json:"coin"`
	Address string  `json:"address"`
	Webhook string  `json:"webhook"`
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64) error
}

type Storage interface {
	Tracker
	Lookup(coin uint, addresses ...string) ([]Subscription, error)
	Add(Subscription) error
	Remove(coin uint, address string) error
	Contains(coin uint, address string) (bool, error)
}

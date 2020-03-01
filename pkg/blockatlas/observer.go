package blockatlas

import "strconv"

type Subscriptions map[string][]string

type SubscriptionOperation string

type SubscriptionEvent struct {
	Subscriptions Subscriptions         `json:"subscriptions"`
	GUID          string                `json:"guid"`
	Operation     SubscriptionOperation `json:"operation"`
}

type CoinStatus struct {
	Height int64  `json:"height"`
	Error  string `json:"error,omitempty"`
}

type Observer struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (e *SubscriptionEvent) ParseSubscriptions() []Subscription {
	subs := make([]Subscription, 0)
	for coinStr, perCoin := range e.Subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, Subscription{
				Coin:    uint(coin),
				Address: addr,
				Webhook: e.GUID,
			})
		}
	}
	return subs
}

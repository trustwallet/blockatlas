package blockatlas

import "strconv"

type Subscriptions map[string][]string

type SubscriptionOperation string

type SubscriptionEvent struct {
	NewSubscriptions Subscriptions         `json:"new_subscriptions"`
	OldSubscriptions Subscriptions         `json:"old_subscriptions"`
	GUID             string                `json:"guid"`
	Operation        SubscriptionOperation `json:"operation"`
}

type CoinStatus struct {
	Height int64  `json:"height"`
	Error  string `json:"error,omitempty"`
}

type Observer struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (e *SubscriptionEvent) ParseSubscriptions(s Subscriptions) []Subscription {
	subs := make([]Subscription, 0)
	for coinStr, perCoin := range s {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, Subscription{
				Coin:    uint(coin),
				Address: addr,
				GUID:    e.GUID,
			})
		}
	}
	return subs
}

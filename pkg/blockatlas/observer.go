package blockatlas

import "strconv"

type (
	Subscriptions map[string][]string

	SubscriptionOperation string

	SubscriptionEvent struct {
		Subscriptions Subscriptions         `json:"subscriptions"`
		Id            uint                  `json:"id"`
		Operation     SubscriptionOperation `json:"operation"`
	}

	Subscription struct {
		Coin    uint   `json:"coin"`
		Address string `json:"address"`
		Id      uint   `json:"id"`
	}

	CoinStatus struct {
		Height int64  `json:"height"`
		Error  string `json:"error,omitempty"`
	}

	Observer struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
)

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
				Id:      e.Id,
			})
		}
	}
	return subs
}

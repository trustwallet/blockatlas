package blockatlas

import "strconv"

type (
	Subscriptions map[string][]string

	SubscriptionOperation string

	SubscriptionEvent struct {
		Subscriptions Subscriptions         `json:"subscriptions"`
		Operation     SubscriptionOperation `json:"operation"`
	}

	Subscription struct {
		Coin    uint   `json:"coin"`
		Address string `json:"address"`
	}
)

func (v *Subscription) AddressID() string {
	return GetAddressID(strconv.Itoa(int(v.Coin)), v.Address)
}

func GetAddressID(coin, address string) string {
	return coin + "_" + address
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
			})
		}
	}
	return subs
}

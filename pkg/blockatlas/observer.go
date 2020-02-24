package blockatlas

import "strconv"

type Subscriptions map[string][]string

type Webhook struct {
	Subscriptions Subscriptions `json:"subscriptions"`
	Webhook       string        `json:"webhook"`
}

type CoinStatus struct {
	Height int64  `json:"height"`
	Error  string `json:"error,omitempty"`
}

type Observer struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (w *Webhook) ParseSubscriptions() []Subscription {
	subs := make([]Subscription, 0)
	for coinStr, perCoin := range w.Subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, Subscription{
				Coin:    uint(coin),
				Address: addr,
				Webhook: w.Webhook,
			})
		}
	}
	return subs
}

package models

type Observer struct{
	Coin    uint    `json:"coin"`
	Address string  `json:"address"`
	Webhook string  `json:"webhook"`
}

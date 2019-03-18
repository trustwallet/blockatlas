package models

type OrderTx struct {
	Kind         string `json:"kind"`
	Id           string `json:"id"`
	Fee          string `json:"fee"`
	SellingAsset string `json:"selling_asset"`
	SellingValue string `json:"selling_amount"`
	BuyingAsset  string `json:"buying_asset"`
	BuyingValue  string `json:"buying_value"`
}

func (_ *OrderTx) Type() string {
	return TxOrder
}

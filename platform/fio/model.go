package fio

type GetPubAddressRequest struct {
	FioAddress string `json:"fio_address"`
	TokenCode  string `json:"token_code"`
}

type GetPubAddressResponse struct {
	PublicAddress string `json:"public_address"`
	BlockNum      int    `json:"block_num"`
	Message       string `json:"message"`
}

package fio

// GetPubAddressRequest request struct for get_pub_address
type GetPubAddressRequest struct {
	FioAddress string `json:"fio_address"`
	TokenCode  string `json:"token_code"`
	ChainCode  string `json:"chain_code"`
}

// GetPubAddressResponse response struct for get_pub_address
type GetPubAddressResponse struct {
	PublicAddress string `json:"public_address"`
	BlockNum      int    `json:"block_num"`
	Message       string `json:"message"`
}

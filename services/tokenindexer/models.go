package tokenindexer

type (
	Request struct {
		From int64
	}

	GetTokensByAddressRequest struct {
		AddressesByCoin map[string][]string `json:"addresses"`
		From            uint                `json:"from"`
	}

	GetTokensByAddressResponse []string
)

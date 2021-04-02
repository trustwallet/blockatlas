package tokenindexer

type (
	Request struct {
		From int64
	}

	GetTokensByAddressRequest struct {
		AddressesByCoin map[string][]string `json:"addresses"`
		From            uint                `json:"from"`
	}

	GetTokensAsset struct {
		AssetId   string `json:"asset_id"`
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
	}

	GetTokensByAddressResponse []GetTokensAsset
)

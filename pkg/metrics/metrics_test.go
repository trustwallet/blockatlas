package metrics

import "testing"

func Test_removeSensitiveInfo(t *testing.T) {
	tests := []struct {
		name string
		info string
		want string
	}{
		{"Remove Nimiq address", "/v1/nimiq/NQ43 J7G6 K6T8 H5KJ 5CXN Q5JK 2GJ4 6DSB 7PUH", "/v1/nimiq/"},
		{"Remove Tezos address", "/v1/tezos/tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q", "/v1/tezos/"},
		{"Remove Tron info", "https://api.trongrid.io/v1/accounts/TPJYCz8ppZNyvw7pTwmjajcx4Kk1MmEUhD/transactions?limit=200&only_confirmed=true&token_id=1000011", "https://api.trongrid.io/v1/accounts//transactions?limit&only_confirmed&token_id"},
		{"Remove asset id", "https://api.trongrid.io/v1/assets/1000570?", "https://api.trongrid.io/v1/assets/?"},
		{"Remove collection id", "/v2/ethereum/collections//collection/---enjin-old", "/v2/ethereum/collections//collection/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSensitiveInfo(tt.info); got != tt.want {
				t.Errorf("removeSensitiveInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

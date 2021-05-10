package oasis

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
	"testing"
)

func TestNormalizeTx(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantTx   types.Tx
	}{
		{
			name:     "Test normalize successful transaction without fee",
			filename: "tx_without_fee.json",
			wantTx: types.Tx{
				ID:       "a49afb8055ef3bbb4fca1e162886ab32b71c5a0d49555793342971237b031972",
				Coin:     coin.OASIS,
				From:     "oasis1qpcgnf84hnvvfvzup542rhc8kjyvqf4aqqlj5kqh",
				To:       "oasis1qz9re9hc0k9qxrhvww7x9zrfv8x8jpr4kcr2twr2",
				Fee:      "0",
				Date:     1610472221000,
				Block:    1502238,
				Status:   types.StatusCompleted,
				Sequence: 5,
				Meta: types.Transfer{
					Value:    "170000000000",
					Symbol:   "ROSE",
					Decimals: 9,
				},
			},
		},
		{
			name:     "Test normalize successful transaction with fee",
			filename: "tx_with_fee.json",
			wantTx: types.Tx{
				ID:       "2b06966eaf27d830cc3c91fe2e38c0d26d38430cf8754be786f66a084ab127d2",
				Coin:     coin.OASIS,
				From:     "oasis1qp29h8ykmxet46eqzw0wennrmmy4al3xzv37m3ca",
				To:       "oasis1qz9re9hc0k9qxrhvww7x9zrfv8x8jpr4kcr2twr2",
				Fee:      "15000",
				Date:     1605717688000,
				Block:    702410,
				Status:   types.StatusCompleted,
				Sequence: 1,
				Meta: types.Transfer{
					Value:    "1000000000",
					Symbol:   "ROSE",
					Decimals: 9,
				},
			},
		},
		{
			name:     "Test normalize transaction with error",
			filename: "tx_with_error.json",
			wantTx: types.Tx{
				ID:       "6df10a2d114739ee6007fbd2ae0905b2ae81be6fc1d8ff6c5a9f404923263b84",
				Coin:     coin.OASIS,
				From:     "oasis1qrmp3lmcuvxhr9dq90mrrxe2yxzwfqw9xcvqujpu",
				To:       "oasis1qqnv3peudzvekhulf8v3ht29z4cthkhy7gkxmph5",
				Fee:      "10",
				Date:     1605774037000,
				Block:    712027,
				Status:   types.StatusError,
				Error:    "insufficient balance",
				Sequence: 1,
				Meta: types.Transfer{
					Value:    "10",
					Symbol:   "ROSE",
					Decimals: 9,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Transaction

			err := mock.JsonModelFromFilePath("mocks/"+tt.filename, &srcTx)
			assert.Nil(t, err)

			gotTx := NormalizeTx(srcTx)
			assert.ObjectsAreEqual(gotTx, tt.wantTx)
		})
	}
}

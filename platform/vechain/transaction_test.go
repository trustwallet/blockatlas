package vechain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

func TestNormalizeTransaction(t *testing.T) {
	var (
		transferSrc, _ = mock.JsonStringFromFilePath("mocks/transfer.json")
		trxId, _       = mock.JsonStringFromFilePath("mocks/tx_id.json")
	)

	tests := []struct {
		name     string
		addr     string
		txData   string
		txId     string
		expected types.Tx
	}{
		{"Test normalize VET transfer transaction", "0xb5e883349e68ab59307d1604555ac890fac47128", transferSrc, trxId, types.Tx{
			ID:        "0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7",
			Coin:      coin.VECHAIN,
			From:      "0xB5e883349e68aB59307d1604555AC890fAC47128",
			To:        "0x2c7A8d5ccE0d5E6a8a31233B7Dc3DAE9AaE4b405",
			Date:      1574410670,
			Type:      types.TxTransfer,
			Fee:       types.Amount("21000"),
			Status:    types.StatusCompleted,
			Block:     4395940,
			Direction: types.DirectionOutgoing,
			Meta: types.Transfer{
				Value:    types.Amount("1347000000000000000"),
				Decimals: 18,
				Symbol:   "VET",
			},
		}},
	}

	platform := Platform{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tx LogTransfer
			err := json.Unmarshal([]byte(tt.txData), &tx)
			assert.Nil(t, err)

			var tId Tx
			errTrxID := json.Unmarshal([]byte(tt.txId), &tId)
			assert.Nil(t, errTrxID)

			actual, err := platform.NormalizeTransaction(tx, tId, tt.addr)
			assert.Nil(t, err)

			assert.Equal(t, tt.expected, actual, "tx don't equal")
		})
	}
}

func TestNormalizeTokenTransaction(t *testing.T) {
	tests := []struct {
		name        string
		txFile      string
		receiptFile string
		address     string
		expected    types.Txs
		wantErr     error
	}{
		{
			name:        "Normalize outgoing VTHO tx",
			txFile:      "outgoing_vtho_tx.json",
			receiptFile: "outgoing_vtho_receipt.json",
			address:     "0xe99399dd211eF54c301A5d1AA813471d92122eA8",
			expected: types.Txs{{
				ID:        "0x0677f91de4787d295087acec0a7ba317b0019fbf296fed630fdb5afbfca97a58",
				Coin:      coin.VECHAIN,
				From:      "0xe99399dd211eF54c301A5d1AA813471d92122eA8",
				To:        "0x0000000000000000000000000000456E65726779",
				Date:      1610958570,
				Type:      types.TxTokenTransfer,
				Fee:       types.Amount("36518000000000000000"),
				Status:    types.StatusCompleted,
				Block:     8045756,
				Direction: types.DirectionOutgoing,
				Meta: types.TokenTransfer{
					Name:     gasTokenName,
					Symbol:   gasTokenSymbol,
					TokenID:  "0x0000000000000000000000000000456E65726779",
					From:     "0xe99399dd211eF54c301A5d1AA813471d92122eA8",
					To:       "0xB5e883349e68aB59307d1604555AC890fAC47128",
					Value:    types.Amount("7000000000000000000"),
					Decimals: 18,
				},
			}},
			wantErr: nil,
		},
		{
			name:        "Normalize incoming VTHO tx",
			txFile:      "incoming_vtho_tx.json",
			receiptFile: "incoming_vtho_receipt.json",
			address:     "0xe99399dd211eF54c301A5d1AA813471d92122eA8",
			expected: types.Txs{{
				ID:        "0xb356fa7b3a371f1518a5f9bc51e951d0dac2ef04d58b532c7ca50a52aa5cddb4",
				Coin:      coin.VECHAIN,
				From:      "0xB5e883349e68aB59307d1604555AC890fAC47128",
				To:        "0x0000000000000000000000000000456E65726779",
				Date:      1610958460,
				Type:      types.TxTokenTransfer,
				Fee:       types.Amount("36582000000000000000"),
				Status:    types.StatusCompleted,
				Block:     8045745,
				Direction: types.DirectionIncoming,
				Meta: types.TokenTransfer{
					Name:     gasTokenName,
					Symbol:   gasTokenSymbol,
					TokenID:  "0x0000000000000000000000000000456E65726779",
					From:     "0xB5e883349e68aB59307d1604555AC890fAC47128",
					To:       "0xe99399dd211eF54c301A5d1AA813471d92122eA8",
					Value:    types.Amount("1000000000000000000000"),
					Decimals: 18,
				},
			}},
			wantErr: nil,
		},
		{
			name:        "Normalize reverted token transfer",
			txFile:      "reverted_tx.json",
			receiptFile: "reverted_receipt.json",
			address:     "0x7cFFB7632252Bae3766734d61F148f0Ea78Fc08C",
			expected:    types.Txs{},
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tx Tx
			err := mock.JsonModelFromFilePath("mocks/"+tt.txFile, &tx)
			assert.Nil(t, err)

			var receipt TxReceipt
			err = mock.JsonModelFromFilePath("mocks/"+tt.receiptFile, &receipt)
			assert.Nil(t, err)

			actual, err := NormalizeTokenTransaction(tx, receipt)
			assert.Equal(t, err, tt.wantErr)

			if len(actual) != 0 {
				updateTransactionDirection(&actual[0], tt.address)
			}
			assert.Equal(t, tt.expected, actual, "tx don't equal")
		})
	}
}

func Test_hexToInt(t *testing.T) {
	tests := []struct {
		name     string
		hex      string
		expected uint64
		wantErr  bool
	}{
		{"value 1", "0x603ca6b1879375dc", 6934600807657731548, false},
		{"value 2", "0x38d7ea4c68000", 1000000000000000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hexToInt(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("hexToInt() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_getRecipientAddress(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want string
	}{
		{"hex 1", "0x000000000000000000000000b5e883349e68ab59307d1604555ac890fac47128", "0xb5e883349e68ab59307d1604555ac890fac47128"},
		{"hex 2", "0x000000000000000000000000f3586684107ce0859c44aa2b2e0fb8cd8731a15a", "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRecipientAddress(tt.hex); got != tt.want {
				t.Errorf("getRecipientAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTransactionDirection(t *testing.T) {
	addr1 := "0xb5e883349e68ab59307d1604555ac890fac47128"
	addr2 := "0eec2bbedbb8b18357dab0b753cd1893bb832284"
	tests := []struct {
		name      string
		sender    string
		recipient string
		address   string
		expected  types.Direction
		expectErr bool
	}{
		{"Self direction for addr1", addr1, addr1, addr1, types.DirectionSelf, false},
		{"Self direction for addr2", addr2, addr2, addr2, types.DirectionSelf, false},
		{"Out direction", addr1, addr2, addr1, types.DirectionOutgoing, false},
		{"In direction", addr1, addr2, addr2, types.DirectionIncoming, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := getTransactionDirection(tt.sender, tt.recipient, tt.address)
			if tt.expectErr {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

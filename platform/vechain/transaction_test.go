package vechain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	transferSrc, _     = mock.JsonStringFromFilePath("mocks/transfer.json")
	trxId, _           = mock.JsonStringFromFilePath("mocks/tx_id.json")
	transferLogSrc, _  = mock.JsonStringFromFilePath("mocks/transfer_log.json")
	trxReceipt, _      = mock.JsonStringFromFilePath("mocks/transfer_receipt.json")
	revertedTx, _      = mock.JsonStringFromFilePath("mocks/reverted_tx.json")
	revertedReceipt, _ = mock.JsonStringFromFilePath("mocks/reverted_receipt.json")
)

func TestNormalizeTransaction(t *testing.T) {
	tests := []struct {
		name     string
		addr     string
		txData   string
		txId     string
		expected types.Tx
	}{
		{"Test normalize VET transfer transaction", "0xb5e883349e68ab59307d1604555ac890fac47128", transferSrc, trxId, types.Tx{
			ID:        "0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7",
			Coin:      coin.VET,
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
		name      string
		txData    string
		txReceipt string
		expected  types.TxPage
	}{
		{"Normalize VIP180 token transfer", transferLogSrc, trxReceipt, types.TxPage{
			{
				ID:        "0x42f5eba46ddcc458243c753545a3faa849502d078efbc5b74baddea9e6ea5b04",
				Coin:      coin.VET,
				From:      "0x2c7A8d5ccE0d5E6a8a31233B7Dc3DAE9AaE4b405",
				To:        "0x0000000000000000000000000000456E65726779",
				Date:      1574278180,
				Type:      types.TxTokenTransfer,
				Fee:       types.Amount("36582000000000000000"),
				Status:    types.StatusCompleted,
				Block:     4382764,
				Direction: types.DirectionIncoming,
				Meta: types.TokenTransfer{
					Name:     gasTokenName,
					Symbol:   gasTokenSymbol,
					TokenID:  "0x0000000000000000000000000000456E65726779",
					From:     "0x2c7A8d5ccE0d5E6a8a31233B7Dc3DAE9AaE4b405",
					To:       "0xB5e883349e68aB59307d1604555AC890fAC47128",
					Value:    types.Amount("68000000000000000000"),
					Decimals: 18,
				},
			},
		}},
		{"Normalize reverted token transfer", revertedTx, revertedReceipt, types.TxPage{
			{
				ID:     "0x7fae32a743e42eaec54642e2a5742a185299f5b4bedaf12c60f65705661de932",
				Coin:   coin.VET,
				From:   "0x7cFFB7632252Bae3766734d61F148f0Ea78Fc08C",
				To:     "0xf8e1fAa0367298b55F57Ed17F7a2FF3F5F1D1628",
				Date:   1610326580,
				Type:   types.TxTokenTransfer,
				Fee:    types.Amount("82618000000000000000"),
				Status: types.StatusError,
				Block:  7982675,
			},
		}},
	}

	platform := Platform{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tx Tx
			err := json.Unmarshal([]byte(tt.txData), &tx)
			assert.Nil(t, err)

			var receipt TxReceipt
			errR := json.Unmarshal([]byte(tt.txReceipt), &receipt)
			assert.Nil(t, errR)

			actual, err := platform.NormalizeTokenTransaction(tx, receipt)
			assert.Nil(t, err)

			assert.Equal(t, len(actual), 1, "tx could not be normalized")
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

func Test_getTokenTransactionDirectory(t *testing.T) {
	addr1 := "0xb5e883349e68ab59307d1604555ac890fac47128"
	addr2 := "0eec2bbedbb8b18357dab0b753cd1893bb832284"
	tests := []struct {
		name         string
		originSender string
		topicsFrom   string
		topicsTo     string
		expected     types.Direction
		expectErr    bool
	}{
		{"Self direction", addr1, addr1, addr1, types.DirectionSelf, false},
		{"In direction", addr1, addr1, addr2, types.DirectionIncoming, false},
		{"Out direction", addr1, addr2, addr1, types.DirectionOutgoing, false},
		{"Unknown direction", addr1, addr2, addr2, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := getTokenTransactionDirectory(tt.originSender, tt.topicsFrom, tt.topicsTo)
			if tt.expectErr {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_getTransferDirectory(t *testing.T) {
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
			actual, err := getTransferDirectory(tt.sender, tt.recipient, tt.address)
			if tt.expectErr {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

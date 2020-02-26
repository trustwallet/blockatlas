package vechain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `{
        "sender": "0xb5e883349e68ab59307d1604555ac890fac47128",
        "recipient": "0x2c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405",
        "amount": "0x12b1815d00738000",
        "meta": {
            "blockID": "0x004313a4bd4286e821b684cc1749deb3df12fa2a8114435fbd35baa155e82016",
            "blockNumber": 4395940,
            "blockTimestamp": 1574410670,
            "txID": "0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7",
            "txOrigin": "0xb5e883349e68ab59307d1604555ac890fac47128",
            "clauseIndex": 0
        }
    }`
const trxId = `{
	"gas": 21000,
	"nonce": "0x8cff29df64a414f8"
}`

var expectedTransfer = blockatlas.Tx{
	ID:        "0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7",
	Coin:      coin.VET,
	From:      "0xB5e883349e68aB59307d1604555AC890fAC47128",
	To:        "0x2c7A8d5ccE0d5E6a8a31233B7Dc3DAE9AaE4b405",
	Date:      1574410670,
	Type:      blockatlas.TxTransfer,
	Fee:       blockatlas.Amount("21000"),
	Status:    blockatlas.StatusCompleted,
	Block:     4395940,
	Direction: blockatlas.DirectionOutgoing,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1347000000000000000"),
		Decimals: 18,
		Symbol:   "VET",
	},
}

func TestNormalizeTransaction(t *testing.T) {
	tests := []struct {
		name     string
		addr     string
		txData   string
		txId     string
		expected blockatlas.Tx
	}{
		{"Test normalize VET transfer transaction", "0xb5e883349e68ab59307d1604555ac890fac47128", transferSrc, trxId, expectedTransfer},
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

const transferLogSrc = `{
    "id": "0x42f5eba46ddcc458243c753545a3faa849502d078efbc5b74baddea9e6ea5b04",
    "chainTag": 74,
    "blockRef": "0x0042e02a2ae04200",
    "expiration": 720,
    "clauses": [
        {
            "to": "0x0000000000000000000000000000456e65726779",
            "value": "0x0",
            "data": "0xa9059cbb000000000000000000000000b5e883349e68ab59307d1604555ac890fac47128000000000000000000000000000000000000000000000003afb087b876900000"
        }
    ],
    "gasPriceCoef": 0,
    "gas": 80000,
    "origin": "0x2c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405",
    "delegator": null,
    "nonce": "0x4a8569d",
    "dependsOn": null,
    "size": 189,
    "meta": {
        "blockID": "0x0042e02cebd1bec003d31526dba338c1b9eeeefdef722fb147e9d31690fbff1e",
        "blockNumber": 4382764,
        "blockTimestamp": 1574278180
    }
}`

const trxReceipt = `{
    "gasUsed": 36582,
    "gasPayer": "0x2c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405",
    "paid": "0x1fbad5f2e25570000",
    "reward": "0x984d9c8dd8008000",
    "reverted": false,
    "meta": {
        "blockID": "0x0042e02cebd1bec003d31526dba338c1b9eeeefdef722fb147e9d31690fbff1e",
        "blockNumber": 4382764,
        "blockTimestamp": 1574278180,
        "txID": "0x42f5eba46ddcc458243c753545a3faa849502d078efbc5b74baddea9e6ea5b04",
        "txOrigin": "0x2c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405"
    },
    "outputs": [
        {
            "contractAddress": null,
            "events": [
                {
                    "address": "0x0000000000000000000000000000456e65726779",
                    "topics": [
                        "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                        "0x0000000000000000000000002c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405",
                        "0x000000000000000000000000b5e883349e68ab59307d1604555ac890fac47128"
                    ],
                    "data": "0x000000000000000000000000000000000000000000000003afb087b876900000"
                }
            ],
            "transfers": []
        }
    ]
}`

var expectedTransferLog = blockatlas.TxPage{
	{
		ID:        "0x42f5eba46ddcc458243c753545a3faa849502d078efbc5b74baddea9e6ea5b04",
		Coin:      coin.VET,
		From:      "0x2c7A8d5ccE0d5E6a8a31233B7Dc3DAE9AaE4b405",
		To:        "0x0000000000000000000000000000456E65726779",
		Date:      1574278180,
		Type:      blockatlas.TxTokenTransfer,
		Fee:       blockatlas.Amount("36582000000000000000"),
		Status:    blockatlas.StatusCompleted,
		Block:     4382764,
		Direction: blockatlas.DirectionIncoming,
		Meta: blockatlas.TokenTransfer{
			Name:     gasTokenName,
			Symbol:   gasTokenSymbol,
			TokenID:  "0x0000000000000000000000000000456E65726779",
			From:     "0x2c7A8d5ccE0d5E6a8a31233B7Dc3DAE9AaE4b405",
			To:       "0xB5e883349e68aB59307d1604555AC890fAC47128",
			Value:    blockatlas.Amount("68000000000000000000"),
			Decimals: 18,
		},
	},
}

func TestNormalizeTokenTransaction(t *testing.T) {
	tests := []struct {
		name      string
		txData    string
		txReceipt string
		expected  blockatlas.TxPage
	}{
		{"Normalize VIP180 token transfer", transferLogSrc, trxReceipt, expectedTransferLog},
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
		expected     blockatlas.Direction
		expectErr    bool
	}{
		{"Self direction", addr1, addr1, addr1, blockatlas.DirectionSelf, false},
		{"In direction", addr1, addr1, addr2, blockatlas.DirectionIncoming, false},
		{"Out direction", addr1, addr2, addr1, blockatlas.DirectionOutgoing, false},
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
		expected  blockatlas.Direction
		expectErr bool
	}{
		{"Self direction for addr1", addr1, addr1, addr1, blockatlas.DirectionSelf, false},
		{"Self direction for addr2", addr2, addr2, addr2, blockatlas.DirectionSelf, false},
		{"Out direction", addr1, addr2, addr1, blockatlas.DirectionOutgoing, false},
		{"In direction", addr1, addr2, addr2, blockatlas.DirectionIncoming, false},
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

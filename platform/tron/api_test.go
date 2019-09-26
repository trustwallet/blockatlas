package tron

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `
{
	"block_timestamp": 1564797900000,
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"value": {
						"amount": 100666888000000,
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferContract"
			}
		]
	},
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

const tokenTransferSrc = `
{
	"block_timestamp": 1564797900000,
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"value": {
						"amount": 2776267,
						"asset_name": "1002000",
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferAssetContract"
			}
		]
	},
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

var transferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1564797900,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "100666888000000",
		Symbol:   "TRX",
		Decimals: 6,
	},
}

var tokenTransferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1564797900,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.TokenTransfer{
		Name:     "BitTorrent",
		Symbol:   "BTT",
		TokenID:  "1002000",
		Decimals: 6,
		Value:    "2776267",
		From:     "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		To:       "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	},
}

var assetInfo = AssetInfo{Name: "BitTorrent", Symbol: "BTT", Decimals: 6, ID: "1002000"}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    &transferDst,
	})
}

func TestNormalizeTokenTransfer(t *testing.T) {
	testNormalizeTokenTransfer(t, &test{
		name:        "token transfer",
		apiResponse: tokenTransferSrc,
		expected:    &tokenTransferDst,
	})
}

func testNormalizeTokenTransfer(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}
	res, err := NormalizeTokenTransfer(&srcTx, assetInfo)
	if err != nil {
		t.Errorf("%s: tx could not be normalized", _test.name)
		return
	}

	actual, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func testNormalize(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}
	res, ok := Normalize(&srcTx)
	if !ok {
		t.Errorf("%s: tx could not be normalized", _test.name)
		return
	}

	actual, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

var tokenDst = blockatlas.Token{
	Name:     "Test",
	Symbol:   "TST",
	Decimals: 8,
	TokenID:  "1",
	Coin:     195,
	Type:     "TRC10",
}

func TestNormalizeToken(t *testing.T) {
	asset := AssetInfo{Name: "Test", Symbol: "TST", ID: "1", Decimals: 8}
	actual := NormalizeToken(asset)
	assert.Equal(t, tokenDst, actual)
}

func TestNormalizeValidator(t *testing.T) {
	validator := Validator{Address: "414d1ef8673f916debb7e2515a8f3ecaf2611034aa"}

	actual, _ := normalizeValidator(validator)
	expected := blockatlas.Validator{
		ID:     "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp",
		Status: true,
		Reward: blockatlas.StakingReward{
			Annual: Annual,
		},
		LockTime:      259200,
		MinimumAmount: "1000000",
	}
	assert.Equal(t, expected, actual)
}

const delegationsSrc = `
[
  {
    "address": "419241920da7d6bb487a33a6df3838e3d208f0b251",
    "balance": 27075639,
	"frozen": [
	  {
		"expire_time": 1569728532000,
		"frozen_balance": 35000000
	  }
	],
    "trc20": [
      {
        "TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY": "990000800"
      },
      {
        "TLa2f6VPqDgRE67v1736s7bJ8Ray5wYjU7": "11000000"
      },
      {
        "TGbhcodQ1jRWB3ZywmfwsRTh4rwbiL2mzh": "7492500"
      }
    ],
    "votes": [
      {
        "vote_address": "414d1ef8673f916debb7e2515a8f3ecaf2611034aa",
        "vote_count": 21
      },
      {
        "vote_address": "4192c5d96c3b847268f4cb3e33b87ecfc67b5ce3de",
        "vote_count": 5
      },
      {
        "vote_address": "4192c5d96c3b847268f4cb3e33b87ecfc67b5ce3de",
        "vote_count": 5
      }
    ]
  }
]`

func TestNormalizeDelegations(t *testing.T) {
	var accountData []AccountsData
	err := json.Unmarshal([]byte(delegationsSrc), &accountData)
	assert.NoError(t, err)
	assert.NotNil(t, accountData)
	assert.Len(t, accountData, 1)
	assert.Len(t, accountData[0].Votes, 3)
	assert.Len(t, accountData[0].Frozen, 1)

	var tronCoin = coin.Tron()
	expected := []blockatlas.Delegation{
		{
			Delegator: accountData[0].Votes[0].VoteAddress,
			Value:     "21000000",
			Coin:      tronCoin.External(),
			Status:    blockatlas.DelegationStatusPending,
		},
		{
			Delegator: accountData[0].Votes[1].VoteAddress,
			Value:     "5000000",
			Coin:      tronCoin.External(),
			Status:    blockatlas.DelegationStatusPending,
		},
		{
			Delegator: accountData[0].Votes[2].VoteAddress,
			Value:     "5000000",
			Coin:      tronCoin.External(),
			Status:    blockatlas.DelegationStatusPending,
		},
	}
	result := NormalizeDelegations(accountData)
	assert.Equal(t, result, expected)
}

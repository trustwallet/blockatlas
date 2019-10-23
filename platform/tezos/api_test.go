package tezos

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `
{
  "ops": [
    {
      "hash": "opN4YjaBwngT8Csz5gKzdwfm78cNquWwcHkjxfHqqstCAT5HWcM",
      "type": "transaction",
      "block": "BKq8skWbocNHvYw2af2erSvh9UYhPkofrWf1UBxDffhDZHEhUxw",
      "time": "2018-07-04T12:43:27Z",
      "height": 5442,
      "status": "applied",
      "is_success": true,
      "is_contract": false,
      "gas_limit": 200,
      "gas_used": 100,
      "gas_price": 0,
      "volume": 1,
      "fee": 0,
      "sender": "tz1TcgvvzDD4hwHQHdPNGw6ZW9wkomwxaQkP",
      "receiver": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q"
    },
    {
      "hash": "oo3aARP7Y271ZkNi6XqsZZHbzrV1sDwdqD8wrgvSBPaSRK2JDuj",
      "type": "reveal",
      "block": "BL3ET2QcAt7xNU2cnxjrY4iM3Wwe8UHLHCE85rhCiSP8zd26Qnk",
      "time": "2018-07-04T12:50:27Z",
      "height": 5449,
      "status": "applied",
      "is_success": true,
      "is_contract": false,
      "gas_limit": 0,
      "gas_used": 0,
      "gas_price": 0,
      "volume": 0,
      "fee": 0,
      "data": "edpktthB79sCK3xQSekMfuhjHLLC593UW5CHyDR9CueVF68PS1K3ZH",
      "sender": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q"
    },
    {
      "hash": "oo3aARP7Y271ZkNi6XqsZZHbzrV1sDwdqD8wrgvSBPaSRK2JDuj",
      "type": "transaction",
      "block": "BL3ET2QcAt7xNU2cnxjrY4iM3Wwe8UHLHCE85rhCiSP8zd26Qnk",
      "time": "2018-07-04T12:50:27Z",
      "height": 5449,
      "status": "applied",
      "is_success": true,
      "is_contract": false,
      "gas_limit": 360,
      "gas_used": 260,
      "gas_price": 0,
      "volume": 1,
      "fee": 0,
      "sender": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
      "receiver": "tz1TcgvvzDD4hwHQHdPNGw6ZW9wkomwxaQkP"
    }
  ]
}
`

const validatorSrc = `
[
	{"pkh":"tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m","rolls":3726}
]
`

func TestNormalize(t *testing.T) {
	var srcOp Op
	err := json.Unmarshal([]byte(transferSrc), &srcOp)
	assert.NoError(t, err)
	assert.NotNil(t, srcOp)

	expected := []blockatlas.Tx{
		{
			ID:    "opN4YjaBwngT8Csz5gKzdwfm78cNquWwcHkjxfHqqstCAT5HWcM",
			Coin:  coin.XTZ,
			Date:  1530708207,
			From:  "tz1TcgvvzDD4hwHQHdPNGw6ZW9wkomwxaQkP",
			To:    "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
			Fee:   "100",
			Block: 5442,
			Meta: blockatlas.Transfer{
				Value:    blockatlas.Amount("1000000"),
				Symbol:   coin.Coins[coin.XTZ].Symbol,
				Decimals: coin.Coins[coin.XTZ].Decimals,
			},
			Status: "completed",
		},
		{
			ID:    "oo3aARP7Y271ZkNi6XqsZZHbzrV1sDwdqD8wrgvSBPaSRK2JDuj",
			Coin:  coin.XTZ,
			Date:  1530708627,
			From:  "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
			To:    "tz1TcgvvzDD4hwHQHdPNGw6ZW9wkomwxaQkP",
			Fee:   "260",
			Block: 5449,
			Meta: blockatlas.Transfer{
				Value:    blockatlas.Amount("1000000"),
				Symbol:   coin.Coins[coin.XTZ].Symbol,
				Decimals: coin.Coins[coin.XTZ].Decimals,
			},
			Status: "completed",
		},
	}
	result := NormalizeTxs(srcOp.Txs)
	assert.Equal(t, result, expected)
}

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: Annual},
			MinimumAmount: blockatlas.Amount("0"),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
	result := normalizeValidator(v)
	assert.Equal(t, result, expected)
}

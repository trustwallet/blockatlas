package algorand

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const (
	transfer = `
{
  "transactions":[
     {
        "type":"pay",
        "tx":"C2LK3CGBPIGERLPFUXE6INSBJGHOXU7YZMEGELWMVSBASFJYOOQQ",
        "from":"5TSQNIL54GB545B3WLC6OVH653SHAELMHU6MSVNGTUNMOEHAMWG7EC3AA4",
        "fee":1000,
        "first-round":2031300,
        "last-round":2031749,
        "noteb64":"6OZ0TFd0HPw=",
        "round":2031351,
        "poolerror":"",
        "payment":{
           "to":"4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U",
           "close":"",
           "closeamount":0,
           "amount":1,
           "torewards":3237690,
           "closerewards":0
        },
        "fromrewards":0,
        "genesisID":"mainnet-v1.0",
        "genesishashb64":"wGHE2Pwdvd7S12BL5FaOP20EGYesN73ktiC1qzkkit8=",
        "group":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
     }
  ]
}
`
)

var expected = []*blockatlas.Tx{
	{
		ID:     "C2LK3CGBPIGERLPFUXE6INSBJGHOXU7YZMEGELWMVSBASFJYOOQQ",
		Coin:   coin.ALGO,
		From:   "5TSQNIL54GB545B3WLC6OVH653SHAELMHU6MSVNGTUNMOEHAMWG7EC3AA4",
		To:     "4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U",
		Fee:    blockatlas.Amount("1000"),
		Date:   0,
		Block:  2031351,
		Status: blockatlas.StatusCompleted,
		Type:   blockatlas.TxTransfer,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount("1"),
			Symbol:   "ALGO",
			Decimals: 6,
		},
	},
	nil,
	nil,
}

func TestNormalize(t *testing.T) {
	var act TransactionsResponse
	assert.NoError(t, json.Unmarshal([]byte(transfer), &act))
	assert.Equal(t, 1, len(act.Transactions))

	for i, v := range act.Transactions {
		tx, ok := Normalize(v)
		assert.True(t, ok)
		assert.Equal(t, expected[i], &tx)
	}
}

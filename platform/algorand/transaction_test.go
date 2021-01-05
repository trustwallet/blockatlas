package algorand

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

const (
	transfer = `
{
   "transactions":[
      {
         "close-rewards":0,
         "closing-amount":0,
         "confirmed-round":2031351,
         "fee":1000,
         "first-valid":2031300,
         "genesis-hash":"wGHE2Pwdvd7S12BL5FaOP20EGYesN73ktiC1qzkkit8=",
         "genesis-id":"mainnet-v1.0",
         "id":"C2LK3CGBPIGERLPFUXE6INSBJGHOXU7YZMEGELWMVSBASFJYOOQQ",
         "intra-round-offset":57,
         "last-valid":2031749,
         "note":"6OZ0TFd0HPw=",
         "payment-transaction":{
            "amount":1,
            "close-amount":0,
            "receiver":"4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U"
         },
         "receiver-rewards":3237690,
         "round-time":1569123058,
         "sender":"5TSQNIL54GB545B3WLC6OVH653SHAELMHU6MSVNGTUNMOEHAMWG7EC3AA4",
         "sender-rewards":0,
         "signature":{
            "sig":"J1G/vapWXJJjuFcsUPut9ffHrFnXsg1GRQlLyqhTOC0V78zCw3OIAYgeg6k/xiX5NDLLrgy4aYF1hhsEXGZ2Dg=="
         },
         "tx-type":"pay"
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
		Date:   1569123058,
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

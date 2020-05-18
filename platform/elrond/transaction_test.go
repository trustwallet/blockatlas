package elrond

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const txTransferSrc1 = `
{
	"hash":"30d404cc7a42b0158b95f6adfbf9a517627d60f6c7e497c1442dfdb6460285df",
	"nonce":0,
	"round":35462,
	"value":"82516976060558456822",
	"receiver":"erd10yagg2vme2jns9zqf9xn8kl86fkc6dr063vnuj0mz2kk2jw0qwuqmfmaw0",
	"sender":"4294967295",
	"data":"ok",
	"signature":"",
	"timestamp":1587715632,
	"status":"Success",
	"fee": "1000"
}`

const txTransferSrc2 = `
{
	"hash":"30d404cc7a42b0158b95f6adfbf9a517627d60f6c7e497c1442dfdb6460285df",
	"nonce":1,
	"round":100,
	"value":"2000",
	"receiver":"erd1qqqqqqqqqqqqqpgq9lqq6w7tddk3gx0c6qefmfph56ve7styx7ys8j3g0s",
	"sender":"erd1v0ce6rapup6rwma5sltyv05xhp33u543nex75a7j39vsz9m6squq6mxm7y",
	"data":"money",
	"signature":"",
	"timestamp":1588757256,
	"status":"Pending",
	"fee": "1500"
}`

var txTransfer1Normalized = blockatlas.Tx{
	ID:       "30d404cc7a42b0158b95f6adfbf9a517627d60f6c7e497c1442dfdb6460285df",
	Coin:     coin.ERD,
	Date:     int64(1587715632),
	From:     "metachain",
	To:       "erd10yagg2vme2jns9zqf9xn8kl86fkc6dr063vnuj0mz2kk2jw0qwuqmfmaw0",
	Fee:      "1000",
	Status:   blockatlas.StatusCompleted,
	Memo:     "ok",
	Sequence: 0,
	Meta: blockatlas.Transfer{
		Value:    "82516976060558456822",
		Symbol:   coin.Elrond().Symbol,
		Decimals: coin.Elrond().Decimals,
	},
	Direction: blockatlas.DirectionOutgoing,
}

var txTransfer2Normalized = blockatlas.Tx{
	ID:       "30d404cc7a42b0158b95f6adfbf9a517627d60f6c7e497c1442dfdb6460285df",
	Coin:     coin.ERD,
	Date:     int64(1588757256),
	From:     "erd1v0ce6rapup6rwma5sltyv05xhp33u543nex75a7j39vsz9m6squq6mxm7y",
	To:       "erd1qqqqqqqqqqqqqpgq9lqq6w7tddk3gx0c6qefmfph56ve7styx7ys8j3g0s",
	Fee:      "1500",
	Status:   blockatlas.StatusPending,
	Memo:     "money",
	Sequence: 1,
	Meta: blockatlas.Transfer{
		Value:    "2000",
		Symbol:   coin.Elrond().Symbol,
		Decimals: coin.Elrond().Decimals,
	},
	Direction: blockatlas.DirectionOutgoing,
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transferSuccess",
		apiResponse: txTransferSrc1,
		expected:    &txTransfer1Normalized,
	})

	testNormalize(t, &test{
		name:        "transferPending",
		apiResponse: txTransferSrc2,
		expected:    &txTransfer2Normalized,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var tx Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &tx)
	if err != nil {
		t.Error(err)
		return
	}

	normalizedTx, ok := NormalizeTx(tx, tx.Sender)
	require.True(t, ok, _test.name+": cannot normalize tx")

	resJSON, err := json.Marshal(&normalizedTx)
	require.Nil(t, err)

	dstJSON, err := json.Marshal(&_test.expected)
	require.Nil(t, err)

	require.Equal(t, string(dstJSON), string(resJSON))
}

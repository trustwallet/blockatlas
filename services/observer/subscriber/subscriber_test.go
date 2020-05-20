package subscriber

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/tests/integration/setup/testdata"
	"testing"
)

func TestToSubscriptionData(t *testing.T) {
	sub := blockatlas.Subscription{
		Coin:    testdata.EthCoin.ID,
		Address: "A",
		Id:      1,
	}
	sub2 := blockatlas.Subscription{
		Coin:    testdata.EthCoin.ID,
		Address: "B",
		Id:      2,
	}

	expectedModel := models.SubscriptionData{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "A",
	}
	expectedModel1 := models.SubscriptionData{
		SubscriptionId: 2,
		Coin:           &testdata.EthCoin.ID,
		Address:        "B",
	}

	res := ToSubscriptionData([]blockatlas.Subscription{sub, sub2})
	assert.Equal(t, 2, len(res))
	assert.Equal(t, expectedModel, res[0])
	assert.Equal(t, expectedModel1, res[1])
}

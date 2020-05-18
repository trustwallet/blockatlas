package subscriber

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func TestToSubscriptionData(t *testing.T) {
	sub := blockatlas.Subscription{
		Coin:    60,
		Address: "A",
		Id:      1,
	}
	sub2 := blockatlas.Subscription{
		Coin:    60,
		Address: "B",
		Id:      2,
	}

	expectedModel := models.SubscriptionData{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "A",
	}
	expectedModel1 := models.SubscriptionData{
		SubscriptionId: 2,
		Coin:           60,
		Address:        "B",
	}

	res := ToSubscriptionData([]blockatlas.Subscription{sub, sub2})
	assert.Equal(t, 2, len(res))
	assert.Equal(t, expectedModel, res[0])
	assert.Equal(t, expectedModel1, res[1])
}

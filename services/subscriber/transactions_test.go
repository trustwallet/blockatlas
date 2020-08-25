package subscriber

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func TestToSubscriptionData(t *testing.T) {
	sub := blockatlas.Subscription{
		Coin:    60,
		Address: "A",
	}
	sub2 := blockatlas.Subscription{
		Coin:    60,
		Address: "B",
	}

	expected := "60_A"
	expected1 := "60_B"
	res := ToSubscriptionData([]blockatlas.Subscription{sub, sub2})
	assert.Equal(t, 2, len(res))
	assert.Equal(t, expected, res[0])
	assert.Equal(t, expected1, res[1])
}

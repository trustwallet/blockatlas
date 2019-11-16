package cmcmap

import (
	"reflect"
	"testing"
)

func TestCmcMapping_getMap(t *testing.T) {
	tests := []struct {
		name  string
		c     CmcSlice
		wantM CmcMapping
	}{
		{
			"parse mapping 1",
			CmcSlice{{Coin: 3}, {Coin: 10}, {Coin: 44}},
			map[uint]CoinMap{3: {Coin: 3}, 10: {Coin: 10}, 44: {Coin: 44}}},
		{
			"parse mapping 2",
			CmcSlice{{Coin: 3}, {Coin: 10}},
			map[uint]CoinMap{3: {Coin: 3}, 10: {Coin: 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := tt.c.getMap(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("getMap() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

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
			CmcSlice{{Id: 3}, {Id: 10}, {Id: 44}},
			map[uint]CoinMap{3: {Id: 3}, 10: {Id: 10}, 44: {Id: 44}}},
		{
			"parse mapping 2",
			CmcSlice{{Id: 3}, {Id: 10}},
			map[uint]CoinMap{3: {Id: 3}, 10: {Id: 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := tt.c.getMap(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("getMap() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

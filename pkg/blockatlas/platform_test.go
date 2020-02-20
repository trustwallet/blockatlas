package blockatlas

import (
	"reflect"
	"testing"
)

func TestPlatforms_GetPlatformList(t *testing.T) {
	var p Platform
	tests := []struct {
		name string
		ps   Platforms
		want []Platform
	}{
		{
			"test 1",
			Platforms{
				"test1": p,
				"test2": p,
				"test3": p,
			},
			[]Platform{p, p, p},
		}, {
			"test 2",
			Platforms{
				"test1": p,
				"test2": p,
			},
			[]Platform{p, p},
		}, {
			"test 3",
			Platforms{
				"test1": p,
			},
			[]Platform{p},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.GetPlatformList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlatformList() = %v, want %v", got, tt.want)
			}
		})
	}
}

package assets

import (
	"reflect"
	"testing"
)

func TestAssetValidators_toMap(t *testing.T) {
	tests := []struct {
		name string
		av   AssetValidators
		want AssetValidatorMap
	}{
		{"test 1 asset", AssetValidators{{ID: "test1"}}, AssetValidatorMap{"test1": {ID: "test1"}}},
		{"test 2 assets", AssetValidators{{ID: "test1"}, {ID: "test2"}}, AssetValidatorMap{"test1": {ID: "test1"}, "test2": {ID: "test2"}}},
		{"test 3 assets", AssetValidators{{ID: "test1"}, {ID: "test2"}, {ID: "test3"}}, AssetValidatorMap{"test1": {ID: "test1"}, "test2": {ID: "test2"}, "test3": {ID: "test3"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.av.toMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

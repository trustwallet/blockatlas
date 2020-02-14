package assets

import (
	"github.com/stretchr/testify/assert"
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
			got := tt.av.toMap()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAssetValidators_activeValidators(t *testing.T) {
	tests := []struct {
		name string
		av   AssetValidators
		want AssetValidators
	}{
		{
			"test get active validators 1",
			AssetValidators{{ID: "test1", Status: ValidatorStatus{false}}},
			AssetValidators{{ID: "test1", Status: ValidatorStatus{false}}},
		},
		{
			"test get active validators 2",
			AssetValidators{{ID: "test1", Status: ValidatorStatus{true}}},
			AssetValidators{},
		},
		{
			"test get active validators 3",
			AssetValidators{{ID: "test1", Status: ValidatorStatus{true}}, {ID: "test2", Status: ValidatorStatus{false}}},
			AssetValidators{{ID: "test2", Status: ValidatorStatus{false}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.av.activeValidators()
			assert.Equal(t, tt.want, got)
		})
	}
}

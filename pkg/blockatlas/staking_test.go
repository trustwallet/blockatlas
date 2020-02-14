package blockatlas

import (
	"reflect"
	"testing"
)

func TestStakeValidators_ToMap(t *testing.T) {
	tests := []struct {
		name string
		sv   StakeValidators
		want ValidatorMap
	}{
		{"test 1 validator", StakeValidators{{ID: "test1"}}, ValidatorMap{"test1": {ID: "test1"}}},
		{"test 2 validators", StakeValidators{{ID: "test1"}, {ID: "test2"}}, ValidatorMap{"test1": {ID: "test1"}, "test2": {ID: "test2"}}},
		{"test 3 validators", StakeValidators{{ID: "test1"}, {ID: "test2"}, {ID: "test3"}}, ValidatorMap{"test1": {ID: "test1"}, "test2": {ID: "test2"}, "test3": {ID: "test3"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sv.ToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

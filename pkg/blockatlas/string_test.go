package blockatlas

import (
	"fmt"
	"testing"
)

func TestGetValidParameter(t *testing.T) {
	tests := []struct {
		first  string
		second string
		result string
	}{
		{"trust", "wallet", "trust"},
		{"", "wallet", "wallet"},
		{"trust", "", "trust"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("GetValidParameter %d", i), func(t *testing.T) {
			s := GetValidParameter(tt.first, tt.second)
			if s != tt.result {
				t.Errorf("got %q, want %q", s, tt.result)
			}
		})
	}
}

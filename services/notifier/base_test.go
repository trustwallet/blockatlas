package notifier

import (
	"testing"
)

func TestUnprefixedAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    string
		coin    uint
		ok      bool
	}{
		{
			name:    "Test invalid address",
			address: "60",
			want:    "",
			coin:    0,
			ok:      false,
		},
		{
			name:    "Test invalid address 2",
			address: "60_",
			want:    "",
			coin:    0,
			ok:      false,
		},
		{
			name:    "Test invalid coin",
			address: "asdf_0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8",
			want:    "",
			coin:    0,
			ok:      false,
		},
		{
			name:    "Test ETH",
			address: "60_0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8",
			want:    "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8",
			coin:    60,
			ok:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := UnprefixedAddress(tt.address)
			if got != tt.want {
				t.Errorf("UnprefixedAddress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.coin {
				t.Errorf("UnprefixedAddress() got1 = %v, want %v", got1, tt.coin)
			}
			if got2 != tt.ok {
				t.Errorf("UnprefixedAddress() got2 = %v, want %v", got2, tt.ok)
			}
		})
	}
}

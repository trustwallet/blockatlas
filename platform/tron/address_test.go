package tron

import "testing"

func TestHexToAddress(t *testing.T) {
	tests := []struct {
		name    string
		hexAddr string
		want    string
		wantErr bool
	}{
		{
			name:    "Test hex to base58 address",
			hexAddr: "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
			want:    "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB58, err := HexToAddress(tt.hexAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB58 != tt.want {
				t.Errorf("HexToAddress() = %v, want %v", gotB58, tt.want)
			}
		})
	}
}

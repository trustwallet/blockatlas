package bounce

import (
	"testing"
)

func Test_normalizeUrl(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want string
	}{
		{
			name: "Test pancake bunny token uri",
			uri:  "ipfs://QmYu9WwPNKNSZQiTCDfRk7aCR472GURavR9M1qosDmqpev/swapsies.json",
			want: "https://ipfs.io/ipfs/QmYu9WwPNKNSZQiTCDfRk7aCR472GURavR9M1qosDmqpev/swapsies.json",
		},
		{
			name: "Test url with ipfs prefix",
			uri:  "ipfs://ipfs/QmS3hmJqpHpvnCocqv9FTZbcSGDnvuFv4qWY3qnwkMpB9x",
			want: "https://ipfs.io/ipfs/QmS3hmJqpHpvnCocqv9FTZbcSGDnvuFv4qWY3qnwkMpB9x",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeUrl(tt.uri); got != tt.want {
				t.Errorf("normalizeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

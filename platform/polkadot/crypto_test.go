package polkadot

import (
	"encoding/hex"
	"testing"
)

func TestPublicKeyToAddress(t *testing.T) {
	type args struct {
		str     string
		network byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "PublicKey bytes to Kusama address",
			args: args{
				str:     "e8e1b8de72651640e302b62dad1f643ec8b65a3647a7409b2896634db599ed60",
				network: NetworkByteMap["KSM"],
			},
			want: "HqfgRXDgCQcV8KAuTAPGuA1r91iEzinmmNBPkR9kiKhifJq",
		},
		{
			name: "PublicKey bytes to Polkadot address",
			args: args{
				str:     "53d82211c4aadb8c67e1930caef2058a93bc29d7af86bf587fba4aa3b1515037",
				network: NetworkByteMap["DOT"],
			},
			want: "12twBQPiG5yVSf3jQSBkTAKBKqCShQ5fm33KQhH3Hf6VDoKW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := hex.DecodeString(tt.args.str)
			if got := PublicKeyToAddress(bytes, tt.args.network); got != tt.want {
				t.Errorf("PublicKeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

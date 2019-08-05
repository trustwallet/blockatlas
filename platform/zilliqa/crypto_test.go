package zilliqa

import (
	"encoding/hex"
	"testing"
)

func Test_EncodePublicKeyToAddress(t *testing.T) {
	type args struct {
		hexString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encode public key to zil address",
			args: args{
				hexString: "029d25b68a18442590e113132a34bb524695c4291d2c49abf2e4cdd7d98db862c3",
			},
			want: "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodePublicKeyToAddress(tt.args.hexString); got != tt.want {
				t.Errorf("EncodePublicKeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_EncodeKeyHashToAddress(t *testing.T) {
	keyHash, _ := hex.DecodeString("7FCcaCf066a5F26Ee3AFfc2ED1FA9810Deaa632C")
	type args struct {
		keyHash []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encode key hash to zil address",
			args: args{
				keyHash: keyHash,
			},
			want: "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeKeyHashToAddress(tt.args.keyHash); got != tt.want {
				t.Errorf("EncodeKeyHashToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

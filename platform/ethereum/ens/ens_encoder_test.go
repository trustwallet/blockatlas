package ens

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_encodeResolver(t *testing.T) {
	tests := []struct {
		name string
		node string
		want string
	}{
		{
			"Test ohmyname.eth resolver",
			"5ddd0923ace8fe255c0971f8e60d7cd400ae734142a13c14d29a87deb87cdac6",
			"0178b8bf5ddd0923ace8fe255c0971f8e60d7cd400ae734142a13c14d29a87deb87cdac6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, _ := hex.DecodeString(tt.node)
			want, _ := hex.DecodeString(tt.want)
			if got := encodeResolver(node[:]); !reflect.DeepEqual(got, want) {
				t.Errorf("encodeResolver() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_encodeAddr(t *testing.T) {
	type args struct {
		node     string
		coinType uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test encodeAddr",
			args{
				"5ddd0923ace8fe255c0971f8e60d7cd400ae734142a13c14d29a87deb87cdac6",
				60,
			},
			"f1cb7e065ddd0923ace8fe255c0971f8e60d7cd400ae734142a13c14d29a87deb87cdac6000000000000000000000000000000000000000000000000000000000000003c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, _ := hex.DecodeString(tt.args.node)
			want, _ := hex.DecodeString(tt.want)
			if got := encodeAddr(node, tt.args.coinType); !reflect.DeepEqual(got, want) {
				t.Errorf("encodeAddr() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_encodeSupportsInterface(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			"Test encodeSupportsInterface",
			"3b3b57de",
			"01ffc9a73b3b57de",
		},
		{
			"Test encodeSupportsInterface",
			"f1cb7e06",
			"01ffc9a7f1cb7e06",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, _ := hex.DecodeString(tt.id)
			want, _ := hex.DecodeString(tt.want)
			if got := encodeSupportsInterface(id); !reflect.DeepEqual(got, want) {
				t.Errorf("encodeSupportsInterface() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_encodeLegacyAddr(t *testing.T) {
	tests := []struct {
		name string
		node string
		want string
	}{
		{
			"Test encodeLegacyAddr",
			"5ddd0923ace8fe255c0971f8e60d7cd400ae734142a13c14d29a87deb87cdac6",
			"3b3b57de5ddd0923ace8fe255c0971f8e60d7cd400ae734142a13c14d29a87deb87cdac6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, _ := hex.DecodeString(tt.node)
			want, _ := hex.DecodeString(tt.want)
			if got := encodeLegacyAddr(node); !reflect.DeepEqual(got, want) {
				t.Errorf("encodeLegacyAddr() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_encodeFunc(t *testing.T) {
	tests := []struct {
		name string
		fn   string
		want string
	}{
		{
			"Test resolver",
			"resolver(bytes32)",
			"0178b8bf",
		},
		{
			"Test addr",
			"addr(bytes32,uint256)",
			"f1cb7e06",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, _ := hex.DecodeString(tt.want)
			if got := encodeFunc(tt.fn); !reflect.DeepEqual(got, want) {
				t.Errorf("encodeFunc() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_encodeCoinType(t *testing.T) {
	tests := []struct {
		name string
		coin uint64
		want string
	}{
		{
			"Test Ethereum",
			uint64(60),
			"000000000000000000000000000000000000000000000000000000000000003c",
		},
		{
			"Test Bitcoin",
			uint64(0),
			"0000000000000000000000000000000000000000000000000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, _ := hex.DecodeString(tt.want)
			if got := encodeCoinType(tt.coin); !reflect.DeepEqual(got, want) {
				t.Errorf("encodeCoinType() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func Test_decodeBytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes string
		want  string
	}{
		{
			"Test decode bytes",
			"00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000014c36edf48e21cf395b206352a1819de658fd7f988000000000000000000000000",
			"c36edf48e21cf395b206352a1819de658fd7f988",
		},
		{
			"Test decode bytes 2",
			"000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000160014b3df10c6941f4a949f1183281844c3a210cba1e200000000000000000000",
			"0014b3df10c6941f4a949f1183281844c3a210cba1e2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := hex.DecodeString(tt.bytes)
			want, _ := hex.DecodeString(tt.want)
			if got := decodeBytes(bytes); !reflect.DeepEqual(got, want) {
				t.Errorf("decodeBytes() = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

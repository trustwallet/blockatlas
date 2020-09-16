package address

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

func TestEIP55Checksum(t *testing.T) {
	tests := []struct {
		name          string
		unchecksummed string
		want          string
	}{
		{"test checksum 1", "checktest", "0xChecKTeSt"},
		{"test checksum 2", "trustwallet", "0xtrUstWaLlET"},
		{"test checksum number", "16345785d8a0000", "0x16345785d8A0000"},
		{"test checksum hex", "fffdefefed", "0xFfFDEfeFeD"},
		{"test checksum 3", "0x0000000000000000003731342d4f4e452d354639", "0x0000000000000000003731342d4f4E452d354639"},
		{"test checksum 4", "0000000000000000003731342d4f4e452d354639", "0x0000000000000000003731342d4f4E452d354639"},
		{"test checksum Ethereum address", "0x84a0d77c693adabe0ebc48f88b3ffff010577051", "0x84A0d77c693aDAbE0ebc48F88b3fFFF010577051"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EIP55Checksum(tt.unchecksummed); got != tt.want {
				t.Errorf("EIP55Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEIP55ChecksumWanchain(t *testing.T) {
	var (
		addr1Wan      = "0xae96137e0e05681ed2f5d1af272c3ee512939d0f"
		addr1WANEIP55 = "0xaE96137e0E05681Ed2f5d1af272c3EE512939d0f"
		tests         = []struct {
			name          string
			unchecksummed string
			want          string
		}{
			{"test 1", addr1Wan, addr1WANEIP55},
			{"test 2", addr1WANEIP55, addr1WANEIP55},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EIP55ChecksumWanchain(tt.unchecksummed); got != tt.want {
				t.Errorf("EIP55ChecksumWanchain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove0x(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"remove 0x from addres", "0x158079ee67fce2f58472a96584a73c7ab9ac95c1", "158079ee67fce2f58472a96584a73c7ab9ac95c1"},
		{"remove 0x from hash", "0x230798fe22abff459b004675bf827a4089326a296fa4165d0c2ad27688e03e0c", "230798fe22abff459b004675bf827a4089326a296fa4165d0c2ad27688e03e0c"},
		{"remove 0x hex value", "0xfffdefefed", "fffdefefed"},
		{"remove 0x hex number", "0x16345785d8a0000", "16345785d8a0000"},
		{"remove hex without 0x", "trustwallet", "trustwallet"},
		{"remove hex number without 0x", "16345785d8a0000", "16345785d8a0000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove0x(tt.input); got != tt.want {
				t.Errorf("Remove0x() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToAddress(t *testing.T) {
	in := "4182dd6b9966724ae2fdc79b416c7588da67ff1b35"
	expected := "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9"
	got, err := HexToAddress(in)
	if err != nil {
		t.Fatal(err)
	}
	if expected != got {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestToEIP55ByCoinID(t *testing.T) {
	var (
		addr1                        = "0xea674fdde714fd979de3edf0f56aa9716b898ec8"
		addr1EIP55                   = "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"
		wanAddrLowercase             = "0xae96137e0e05681ed2f5d1af272c3ee512939d0f"
		wanAddrEIP55Checksum         = "0xAe96137E0e05681eD2F5D1AF272C3ee512939D0F"
		wanAddrEIP55ChecksumWanchain = "0xaE96137e0E05681Ed2f5d1af272c3EE512939d0f"
		tests                        = []struct {
			name, address, expectedAddress string
			coinID                         uint
		}{
			{"Ethereum", addr1, addr1EIP55, coin.ETH},
			{"Ethereum Classic", addr1, addr1EIP55, coin.ETC},
			{"POA", addr1, addr1EIP55, coin.POA},
			{"Callisto", addr1, addr1EIP55, coin.CLO},
			{"Tomochain", addr1, addr1EIP55, coin.TOMO},
			{"Thunder", addr1, addr1EIP55, coin.TT},
			{"Thunder", addr1, addr1EIP55, coin.TT},
			{"GoChain", addr1, addr1EIP55, coin.GO},
			{"Wanchain 1", wanAddrLowercase, wanAddrEIP55ChecksumWanchain, coin.WAN},
			{"Wanchain 2", wanAddrEIP55Checksum, wanAddrEIP55ChecksumWanchain, coin.WAN},
			{"Non Ethereum like chain 1", "", "", coin.TRX},
			{"Non Ethereum like chain 2", addr1, addr1, coin.BNB},
		}
	)

	t.Run("Test TestToEIP55ByCoinID", func(t *testing.T) {
		for _, tt := range tests {
			actual := ToEIP55ByCoinID(tt.address, tt.coinID)
			assert.Equal(t, tt.expectedAddress, actual)
		}
	})
}

func TestFormatAddress(t *testing.T) {
	var (
		addr1                        = "0xea674fdde714fd979de3edf0f56aa9716b898ec8"
		addr1EIP55                   = "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"
		wanAddrLowercase             = "0xae96137e0e05681ed2f5d1af272c3ee512939d0f"
		wanAddrEIP55Checksum         = "0xAe96137E0e05681eD2F5D1AF272C3ee512939D0F"
		wanAddrEIP55ChecksumWanchain = "0xaE96137e0E05681Ed2f5d1af272c3EE512939d0f"
		tests                        = []struct {
			name, address, expectedAddress string
			coinID                         uint
		}{
			{"Ethereum", addr1, addr1EIP55, coin.ETH},
			{"Ethereum Classic", addr1, addr1EIP55, coin.ETC},
			{"POA", addr1, addr1EIP55, coin.POA},
			{"Callisto", addr1, addr1EIP55, coin.CLO},
			{"Tomochain", addr1, addr1EIP55, coin.TOMO},
			{"Thunder", addr1, addr1EIP55, coin.TT},
			{"Thunder", addr1, addr1EIP55, coin.TT},
			{"GoChain", addr1, addr1EIP55, coin.GO},
			{"Wanchain 1", wanAddrLowercase, wanAddrEIP55ChecksumWanchain, coin.WAN},
			{"Wanchain 2", wanAddrEIP55Checksum, wanAddrEIP55ChecksumWanchain, coin.WAN},
			{"Non Ethereum like chain 1", "", "", coin.TRX},
			{"Non Ethereum like chain 2", addr1, addr1, coin.BNB},
			{"Bitcoin cash case with prefix", "bitcoincash:qzzhnrz43k86r3shen9se96uqeu5mxe0msa2auy85w", "qzzhnrz43k86r3shen9se96uqeu5mxe0msa2auy85w", coin.BCH},
			{"Bitcoin cash case without prefix", "qr5q38d4g02u976jtl7s2ygsewlpaaaylsp2jm6wpf", "qr5q38d4g02u976jtl7s2ygsewlpaaaylsp2jm6wpf", coin.BCH},
			{"Bitcoin cash case without prefix", "qr5q38d4g02u976jtl7s2ygsewlpaaaylsp2jm6wpf", "qr5q38d4g02u976jtl7s2ygsewlpaaaylsp2jm6wpf", coin.BTC},
		}
	)

	t.Run("Test TestToEIP55ByCoinID", func(t *testing.T) {
		for _, tt := range tests {
			actual := FormatAddress(tt.address, tt.coinID)
			assert.Equal(t, tt.expectedAddress, actual)
		}
	})
}

func TestUnprefixedAddress(t *testing.T) {
	address, id, ok := UnprefixedAddress("60_a")
	assert.Equal(t, "a", address)
	assert.Equal(t, uint(60), id)
	assert.True(t, ok)
}

func TestPrefixedAddress(t *testing.T) {
	address := PrefixedAddress(60, "a")
	assert.Equal(t, "60_a", address)
}

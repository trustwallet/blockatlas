package address

import (
	"testing"
)

func TestChecksum(t *testing.T) {
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
	tests := []struct {
		name          string
		unchecksummed string
		want          string
	}{
		{"test checksum 1", "0xae96137e0e05681ed2f5d1af272c3ee512939d0f", "0xaE96137e0E05681Ed2f5d1af272c3EE512939d0f"},
		{"test checksum 2", "0xAe96137E0e05681eD2F5D1AF272C3ee512939D0F", "0xaE96137e0E05681Ed2f5d1af272c3EE512939d0f"},
	}
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

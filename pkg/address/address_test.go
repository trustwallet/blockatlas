package address

import (
	"github.com/stretchr/testify/assert"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EIP55Checksum(tt.unchecksummed); got != tt.want {
				t.Errorf("EIP55Checksum() = %v, want %v", got, tt.want)
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

func TestEncode(t *testing.T) {
	assert.Equal(t, EncodeEIP55(195, "1002000"), "0x0000000000000000003139352D31303032303030")
	assert.Equal(t, EncodeEIP55(714, "AGRI-BD2"), "0x00000000000000003731342d414752492d424432")
	assert.Equal(t, EncodeEIP55(714, "CHZ-ECD"), "0x0000000000000000003731342D43485a2D454344")
	assert.Equal(t, EncodeEIP55(195, "TGbhcodQ1jRWB3ZywmfwsRTh4rwbiL2mzh"), "0x3139352d54476268636f6451316a525742335A79")
	assert.Equal(t, EncodeEIP55(195, "TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY"), "0x3139352d544B54636642454B7070355A5250776D")
}

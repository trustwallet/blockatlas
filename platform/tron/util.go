package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/mr-tron/base58"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

// HexToAddress converts a hex representation of a Tron address
// into a Base58 string with a 4 byte checksum.
func HexToAddress(hexAddr string) (b58 string, err error) {
	bytes, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", errors.E(err, errors.TypePlatformUnmarshal,
			errors.Params{"hexAddr": hexAddr}).PushToSentry()
	}
	var checksum [32]byte
	checksum = sha256.Sum256(bytes)
	checksum = sha256.Sum256(checksum[:])
	bytes = append(bytes, checksum[:4]...)
	b58 = base58.EncodeAlphabet(bytes, base58.BTCAlphabet)
	return
}

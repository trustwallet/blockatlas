package address

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/mr-tron/base58"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"golang.org/x/crypto/sha3"
	"strings"
)

// Decode decodes a hex string with 0x prefix.
func Remove0x(input string) string {
	if strings.HasPrefix(input, "0x") {
		return input[2:]
	}
	return input
}

// Hex returns an EIP55-compliant hex string representation of the address.
func EIP55Checksum(unchecksummed string) string {
	v := []byte(Remove0x(strings.ToLower(unchecksummed)))
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write(v)
	if err != nil {
		logger.Error(err)
	}
	hash := sha.Sum(nil)

	result := v
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	val := string(result)
	return "0x" + val
}

// HexToAddress converts a hex representation of a Tron address
// into a Base58 string with a 4 byte checksum.
func HexToAddress(hexAddr string) (b58 string, err error) {
	bytes, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", errors.E(err, errors.TypePlatformUnmarshal,
			errors.Params{"hexAddr": hexAddr})
	}
	var checksum [32]byte
	checksum = sha256.Sum256(bytes)
	checksum = sha256.Sum256(checksum[:])
	bytes = append(bytes, checksum[:4]...)
	b58 = base58.EncodeAlphabet(bytes, base58.BTCAlphabet)
	return
}

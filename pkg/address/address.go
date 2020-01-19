package address

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

const ethereumAddressLength = 40

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
	sha.Write(v)
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

func EncodeEIP55(coin int, id string) string {
	src := []byte(strconv.Itoa(coin) + "-" + id)
	s := hex.EncodeToString(src)
	count := ethereumAddressLength - len(s)
	if count >= 0 {
		return EIP55Checksum(strings.Repeat("0", count) + s)
	}
	return EIP55Checksum(s[0:ethereumAddressLength])
}

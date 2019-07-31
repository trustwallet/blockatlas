package zilliqa

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
)

const HRP string = "zil"

func EncodePublicKeyToAddress(hexString string) string {
	hexString = strings.TrimPrefix(hexString, "0x")
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return ""
	}
	keyHash := sha256.Sum256(bytes)
	return EncodeKeyHashToAddress(keyHash[12:])
}

func EncodeKeyHashToAddress(keyHash []byte) string {
	conv, err := bech32.ConvertBits(keyHash, 8, 5, true)
	if err != nil {
		return ""
	}
	encoded, err := bech32.Encode(HRP, conv)
	if err != nil {
		return ""
	}
	return encoded
}

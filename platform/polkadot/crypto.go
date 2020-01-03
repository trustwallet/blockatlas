package polkadot

import (
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2b"
)

var ss58Prefix = []byte("SS58PRE")

// PublicKeyToAddress returns an ss58 address string given public key bytes
// see: https://github.com/paritytech/substrate/wiki/External-Address-Format-(SS58)
func PublicKeyToAddress(bytes []byte, network byte) string {
	encode := []byte{network}
	encode = append(encode, bytes...)
	hasher, err := blake2b.New(64, nil)
	if err != nil {
		return ""
	}
	_, err = hasher.Write(append(ss58Prefix, encode...))
	if err != nil {
		return ""
	}
	checksum := hasher.Sum(nil)
	encode = append(encode, checksum[:2]...)
	return base58.Encode(encode)
}

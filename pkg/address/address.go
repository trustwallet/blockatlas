package address

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/mr-tron/base58"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/golibs/coin"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

const prefixBitcoinCash = "bitcoincash:"

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
		log.Error(err)
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
		return "", err
	}
	var checksum [32]byte
	checksum = sha256.Sum256(bytes)
	checksum = sha256.Sum256(checksum[:])
	bytes = append(bytes, checksum[:4]...)
	b58 = base58.EncodeAlphabet(bytes, base58.BTCAlphabet)
	return
}

// Returns an EIP55 Wanchain compliant hex string representation of the address.
// See https://wandevs.org/docs/difference-between-wanchain-and-ethereum/
// https://github.com/wanchain/go-wanchain/blob/b238c203ca7dc6a578d57c0c473bec0fcf2431bd/common/types.go#L173
func EIP55ChecksumWanchain(address string) string {
	v := []byte(Remove0x(strings.ToLower(address)))
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write(v)
	if err != nil {
		log.Error(err)
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
		if result[i] > '9' && hashByte <= 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}

func ToEIP55ByCoinID(str string, coinID uint) string {
	switch coinID {
	case coin.ETH, coin.POA, coin.ETC, coin.TOMO, coin.CLO, coin.TT, coin.GO:
		return EIP55Checksum(str)
	case coin.WAN:
		return EIP55ChecksumWanchain(str)
	default:
		return str
	}
}

func removePrefix(address string) string {
	return strings.TrimPrefix(address, prefixBitcoinCash)
}

func FormatAddress(address string, coinID uint) string {
	switch coinID {
	case coin.ETH, coin.POA, coin.ETC, coin.TOMO, coin.CLO, coin.TT, coin.GO, coin.WAN:
		return ToEIP55ByCoinID(address, coinID)
	case coin.BCH:
		return removePrefix(address)
	default:
		return address
	}
}

func PrefixedAddress(coinID uint, address string) string {
	return strconv.Itoa(int(coinID)) + "_" + address
}

func UnprefixedAddress(address string) (string, uint, bool) {
	result := strings.Split(address, "_")
	if len(result) != 2 {
		return "", 0, false
	}
	id, err := strconv.Atoi(result[0])
	if err != nil {
		return "", 0, false
	}
	return result[1], uint(id), true

}

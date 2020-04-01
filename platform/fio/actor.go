package fio

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func actorFromPublicKeyOrActor(addressOrActor string) string {
	l := len(addressOrActor)
	if l >= 51 && l <= 55 && addressOrActor[:3] == "FIO" {
		// assume public key string
		return actorFromPublicKey(addressOrActor)
	}
	if l <= 13 {
		// assume actor
		return addressOrActor
	}
	return ""
}

func actorFromPublicKey(address string) string {
	pkBytes, err := bytesFromPublicKeyString(address)
	if err != nil {
		return ""
	}
	return actorFromPublicKeyBytes(pkBytes)
}

func actorFromPublicKeyBytes(pkBytes []byte) string {
	shortenedKey := shortenKey(pkBytes)
	name13 := getName(shortenedKey)
	// trim to 12 characters
	return name13[:12]
}

func bytesFromPublicKeyString(address string) ([]byte, error) {
	if address[:3] != "FIO" {
		return nil, errors.E("Invalid FIO public key prefix")
	}
	array := base58.Decode(address[3:])
	if len(array) != 37 {
		return nil, errors.E("Invalid FIO public key length")
	}
	return array, nil
}

func mask12(len int) byte {
	if len == 12 {
		return 0x0f
	}
	return 0x1f
}

func mask0(len int) byte {
	if len == 0 {
		return 0x0f
	}
	return 0x1f
}

func shortenKey(addrKey []byte) uint64 {
	var (
		res uint64 = 0
		i          = 1 // Ignore key head
		l          = 0
	)
	for l <= 12 {
		//assert(i < 33)
		trimmedChar := uint64(addrKey[i] & mask12(l))
		if trimmedChar == 0 {
			i++
			continue
		} // Skip a zero and move to next
		var shuffle byte = 0
		if l < 12 {
			shuffle = byte(5*(12-l) - 1)
		}
		res = res | (trimmedChar << shuffle)
		l++
		i++
	}
	return res
}

func getName(shortKey uint64) string {
	var (
		charmap = ".12345abcdefghijklmnopqrstuvwxyz"
		str     [13]byte
		tmp     = shortKey
	)
	for i := 0; i <= 12; i++ {
		c := charmap[tmp&uint64(mask0(i))]
		str[12-i] = c
		if i == 0 {
			tmp = tmp >> 4
		} else {
			tmp = tmp >> 5
		}
	}
	return string(str[:])
}

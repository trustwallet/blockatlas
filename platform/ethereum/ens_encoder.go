package ethereum

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math/big"

	"golang.org/x/crypto/sha3"
)

func encodeResolver(node []byte) []byte {
	data := make([]byte, 0, 36)
	signature := encodeFunc("resolver(bytes32)")
	data = append(data, signature...)
	data = append(data, node[:]...)
	return data
}

func encodeAddr(node []byte, coinType uint64) []byte {
	data := make([]byte, 0, 68)
	signature := encodeFunc("addr(bytes32,uint256)")
	data = append(data, signature...)
	data = append(data, node...)
	data = append(data, encodeCoinType(coinType)...)
	return data
}

func encodeLegacyAddr(node []byte) []byte {
	data := make([]byte, 0, 36)
	signature := encodeFunc("addr(bytes32)")
	data = append(data, signature...)
	data = append(data, node...)
	return data
}

func encodeFunc(fn string) []byte {
	data := make([]byte, 0, 32)
	sha := sha3.NewLegacyKeccak256()
	if _, err := sha.Write([]byte(fn)); err != nil {
		return data
	}
	sha.Sum(data)
	return data[:4]
}

func encodeCoinType(i uint64) []byte {
	data := make([]byte, 24)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		return data
	}
	data = append(data, buf.Bytes()...)
	return data
}

func decodeHex(hexString string) []byte {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return []byte{}
	}
	return decoded
}

func decodeBytes(b []byte) []byte {
	offset := int64(32)
	count := new(big.Int)
	count.SetBytes(b[:offset])
	length := new(big.Int)
	length.SetBytes(b[offset : offset+count.Int64()])
	offset += count.Int64()
	return b[offset : offset+length.Int64()]
}

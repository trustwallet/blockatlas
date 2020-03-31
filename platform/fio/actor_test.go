package fio

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActorFromPublicKey(t *testing.T) {
	addrArr := []string{
		"FIO6cDpi7vPnvRwMEdXtLnAmFwygaQ8CzD7vqKLBJ2GfgtHBQ4PPy",
		"FIO7uMZoeei5HtXAD24C4yCkpWWbf24bjYtrRNjWdmGCXHZccwuiE",
		"FIO7bxrQUTbQ4mqcoefhWPz1aFieN4fA9RQAiozRz7FrUChHZ7Rb8",
		"FIO6m1fMdTpRkRBnedvYshXCxLFiC5suRU8KDfx8xxtXp2hntxpnf",
		"FIO7Q3XfQ2ocGP1zYst6Sfx5qrsiZ865Cu8So2atrub9JN94so7gt",
		"FIO5kJKNHwctcfUM5XZyiWSqSTM5HTzznJP9F3ZdbhaQAHEVq575o",
		"2odzomo2v4pe", // actor
		"hhq2g4qgycfb", // actor
		"FIO5kJKNHwctcfUM5XZyiWSqSTM5HTzznJP9F3ZdbhaQAHEVq575", // invalid length
		"FIO5kJKNHwctcfUM5XZyiWSqSTM5H",                        // invalid length
		"FIO5kJKNHwct",                                         // assume actor
	}
	actorArr := []string{
		"2odzomo2v4pe",
		"hhq2g4qgycfb",
		"5kmx4qbqlpld",
		"qdfejz2a5wpl",
		"ezsmbcy2opod",
		"ltwagbt4qpuk",
		"2odzomo2v4pe",
		"hhq2g4qgycfb",
		"",
		"",
		"FIO5kJKNHwct",
	}
	for i := range addrArr {
		assert.Equal(t, actorArr[i], actorFromPublicKeyOrActor(actorArr[i]))
	}
}

func TestBytesFromPublicKeyString(t *testing.T) {
	{
		pkBytes, err := bytesFromPublicKeyString("FIO5kJKNHwctcfUM5XZyiWSqSTM5HTzznJP9F3ZdbhaQAHEVq575o")
		assert.Equal(t, nil, err)
		assert.Equal(t, 37, len(pkBytes))
		assert.Equal(t, "0271195c66ec2799e436757a70cd8431d4b17733a097b18a5f7f1b6b085978ff0f343fc54e", hex.EncodeToString(pkBytes))
	}
	{
		pkBytes, err := bytesFromPublicKeyString("FIO6cDpi7vPnvRwMEdXtLnAmFwygaQ8CzD7vqKLBJ2GfgtHBQ4PPy")
		assert.Equal(t, nil, err)
		assert.Equal(t, "02e274495ff4d2f4027bc4d5ead805b0197f19efe526ba2c1c5545ba916d6088a7248a8bf4", hex.EncodeToString(pkBytes))
	}
	{
		pkBytes, err := bytesFromPublicKeyString("FIO6cDpi7vPnvRwMEdXtLnAmFwygaQ8CzD7vqKLBJ2GfgtHBQ4P")
		assert.True(t, err != nil)
		assert.Equal(t, 0, len(pkBytes))
	}
}

func TestShortenKey(t *testing.T) {
	pkBytes, _ := hex.DecodeString("02e274495ff4d2f4027bc4d5ead805b0197f19efe526ba2c1c5545ba916d6088a7248a8bf4")
	assert.Equal(t, uint64(1518832697283783336), shortenKey(pkBytes))
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "2odzomo2v4pec", getName(uint64(1518832697283783336)))
}

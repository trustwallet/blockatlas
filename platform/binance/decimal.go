package binance

import (
	"fmt"
	"strconv"
	"strings"
)

func DecimalToSatoshis(decStr string) (uint64, error) {
	// Split string by comma
	parts := strings.Split(decStr, ".")
	if len(parts) != 2 {
		return 0, fmt.Errorf("not a number")
	}
	leftStr  := parts[0]
	rightStr := parts[1]

	// Find precision
	dot   := strings.IndexRune(decStr, '.')
	exp   := len(decStr) - dot - 1

	// Calculate satoshis by decimal places
	left, err  := strconv.ParseUint(leftStr,  10, 64)
	if err != nil {
		return 0, err
	}
	right, err := strconv.ParseUint(rightStr, 10, 64)
	if err != nil {
		return 0, err
	}

	for i := 0; i < exp; i++ {
		left *= 10
	}
	return left + right, nil
}

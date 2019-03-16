package util

import (
	"fmt"
	"strings"
	"unicode"
)

func DecimalToSatoshis(decStr string) (string, error) {
	out := strings.Replace(decStr, ".", "", 1)
	out = strings.TrimLeft(out, "0")
	for _, c := range out {
		if !unicode.IsNumber(c) {
			return "", fmt.Errorf("not a number")
		}
	}
	return out, nil
}

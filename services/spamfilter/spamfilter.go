package spamfilter

import (
	"strings"
)

var SpamList []string

func ContainsSpam(name string) bool {
	lowerCaseName := strings.ToLower(name)
	for _, word := range SpamList {
		if strings.Contains(lowerCaseName, word) {
			return true
		}
	}
	return false
}

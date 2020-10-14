package spamfilter

import (
	"regexp"
	"strings"
)

var SpamList []string

func ContainsSpam(name string) bool {
	lowerCaseName := strings.ToLower(name)
	for _, word := range SpamList {
		if strings.Contains(lowerCaseName, word) || isURL(lowerCaseName) {
			return true
		}
	}
	return false
}

func isURL(host string) bool {
	var URLRegex = `[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	return regexp.MustCompile(URLRegex).MatchString(host)
}

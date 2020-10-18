package spamfilter

import (
	"regexp"
	"strings"
)

const URLRegex = `[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

var (
	SpamList       []string
	compiledRegexp = regexp.MustCompile(URLRegex)
)

func ContainsSpam(name string) bool {
	if isURL(name) {
		return true
	}
	lowerCaseName := strings.ToLower(name)
	for _, word := range SpamList {
		if strings.Contains(lowerCaseName, word) {
			return true
		}
	}
	return false
}

func isURL(host string) bool {
	return compiledRegexp.MatchString(host)
}

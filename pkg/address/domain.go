package address

import (
	"strings"
)

// Obtain tld from the name, e.g. ".eth" from "nick.eth"
func GetTLD(name, separator string) string {
	lastIdx := strings.LastIndex(name, separator)
	if lastIdx < 0 || lastIdx >= len(name)-1 {
		return ""
	}
	// return tail including separator
	return name[lastIdx:]
}

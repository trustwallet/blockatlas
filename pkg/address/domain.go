package address

import (
	"strings"
)

// Obtain tld from the name, e.g. ".eth" from "nick.eth"
func GetTLD(name, separator string) string {
	lastSeparatorIdx := strings.LastIndex(name, separator)
	if lastSeparatorIdx < 0 || lastSeparatorIdx >= len(name)-1 {
		// no separator inside string
		return ""
	}
	// return tail including separator
	return name[lastSeparatorIdx:]
}

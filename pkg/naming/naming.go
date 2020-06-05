package naming

import (
	"strings"
)

func GetTopDomain(name, separator string) string {
	lastIdx := strings.LastIndex(name, separator)
	if lastIdx < 0 || lastIdx >= len(name)-1 {
		return ""
	}
	// return tail including separator
	return strings.ToLower(name[lastIdx:])
}

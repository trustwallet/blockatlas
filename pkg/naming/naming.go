package naming

import (
	"strings"
)

func GetTLD(name, separator string) string {
	lastIdx := strings.LastIndex(name, separator)
	if lastIdx < 0 || lastIdx >= len(name)-1 {
		return ""
	}
	// return tail including separator
	return name[lastIdx:]
}

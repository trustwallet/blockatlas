package compound

import (
	"encoding/json"
	"sort"
)

// Recursive-sort a json string
func sortJSON(j string) string {
	if len(j) == 0 {
		return j
	}
	if j[0] == '{' {
		// map
		var interfaceMap map[string]interface{}
		if err := json.Unmarshal([]byte(j), &interfaceMap); err != nil {
			return j
		}
		keys := []string{}
		for key := range interfaceMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		// put together
		sorted := "{"
		for idx, key := range keys {
			if idx > 0 {
				sorted += ","
			}
			sorted += "\"" + key + "\":"
			jb, _ := json.Marshal(interfaceMap[key])
			sorted += sortJSON(string(jb))
		}
		sorted += "}"
		return sorted
	}
	if j[0] == '[' {
		// array
		var interfaceArray []interface{}
		if err := json.Unmarshal([]byte(j), &interfaceArray); err != nil {
			return j
		}
		// marshal each element to string
		stringArray := make([]string, len(interfaceArray))
		for idx, val := range interfaceArray {
			jb, _ := json.Marshal(val)
			stringArray[idx] = string(jb)
		}
		// sort elements
		sort.Strings(stringArray)
		// put together
		sorted := "["
		for idx, val := range stringArray {
			if idx > 0 {
				sorted += ","
			}
			sorted += sortJSON(val)
		}
		sorted += "]"
		return sorted
	}
	// non-compound
	return j
}

// Compare two json strings
func CompareJSON(j1, j2 string) bool {
	sorted1 := sortJSON(j1)
	sorted2 := sortJSON(j2)
	return (sorted1 == sorted2)
}

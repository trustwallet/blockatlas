package main

import (
	"testing"
)

func TestSortJSON(t *testing.T) {
	tests := []struct{ name, json1, json2 string }{
		{"simple", "abc", "abc"},
		{"simple num", "123", "123"},
		{"simple quote", `"123"`, `"123"`},
		{"map simple", `{"b": 1, "a": 2, "c": "three"}`, `{"a":2,"b":1,"c":"three"}`},
		{"list simple", `["b", "a", "c"]`, `["a","b","c"]`},
		{"list of maps", `[{"a": "B", "b": 1}, {"a": "A", "b": 2}]`, `[{"a":"A","b":2},{"a":"B","b":1}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortJSON(tt.json1)
			if res != tt.json2 {
				t.Errorf("Wrong result, %v vs %v", res, tt.json2)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestMatchQueryParams(t *testing.T) {
	tests := [][]string{
		{
			"a=1&b=20",
			"a=1&b=20",
			"true",
		},
		{
			"a=1&b=20",
			"a=1&b=20&c=3",
			"true",
		},
		{
			"a=1&b=20",
			"b=20&a=1",
			"true",
		},
		{
			"a=1&b=20",
			"a=1&b=500",
			"false",
		},
		{
			"a=1&b=20",
			"a=123&b=20",
			"false",
		},
		{
			"a=1&b=20",
			"a=1",
			"false",
		},
		{
			"a=1&b=20",
			"b=20",
			"false",
		},
		{
			"",
			"c=500",
			"true",
		},
		{
			"",
			"",
			"true",
		},
	}
	for _, tt := range tests {
		inputExpected := tt[0]
		inputActual := tt[1]
		expectedResult := true
		if tt[2] == "false" {
			expectedResult = false
		}
		result := matchQueryParams(inputExpected, inputActual)
		if result != expectedResult {
			t.Errorf("Did not match, inputExpected %v inputActual %v expectedResult %v", inputExpected, inputActual, expectedResult)
		}
	}
}

package util

import "testing"

func TestDecimalToSatoshis(t *testing.T) {
	assertSatEquals := func(expected string, input string) {
		actual, err := DecimalToSatoshis(input)
		if err != nil {
			t.Error(err)
		}
		if expected != actual {
			t.Errorf("expected: %s, got %s", expected, actual)
		}
	}

	assertSatEquals("10", "1.0")
	assertSatEquals("1", "0.1")
	assertSatEquals("13602", "136.02")
	assertSatEquals("1500000", "0.01500000")
}

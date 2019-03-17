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

func TestDecimalExp(t *testing.T) {
	assertEquals := func(inputDec string, inputExp int, expected string) {
		actual := DecimalExp(inputDec, inputExp)
		if expected != actual {
			t.Errorf("expected: %s * (10^%d) = %s, got %s",
				inputDec, inputExp, expected, actual)
		}
	}

	// No-Op
	assertEquals("0", 300, "0")
	assertEquals("123", 0, "123")
	assertEquals("0.456", 0, "0.456")
	assertEquals("123.456", 0, "123.456")

	// In-Bounds, comma
	assertEquals("12.34", -1, "1.234")
	assertEquals("12.34",  1, "123.4")

	// 1 past bounds, comma
	assertEquals("12.34", -2, "0.1234")
	assertEquals("12.34",  2, "1234")

	// n past bounds, comma
	assertEquals("12.34", -4, "0.001234")
	assertEquals("12.34",  4, "123400")

	// Integer
	assertEquals("1234", -1, "123.4")
	assertEquals("1234",  1, "12340")

	// Denormalized
	assertEquals("0.1234", -1, "0.01234")
	assertEquals("0.1234",  1, "1.234")

	// Tiny
	assertEquals("0.001234", -1, "0.0001234")
	assertEquals("0.001234",  1, "0.01234")
}

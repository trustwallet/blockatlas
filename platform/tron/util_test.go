package tron

import "testing"

func TestHexToAddress(t *testing.T) {
	in := "4182dd6b9966724ae2fdc79b416c7588da67ff1b35"
	expected := "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9"
	got, err := HexToAddress(in)
	if err != nil {
		t.Fatal(err)
	}
	if expected != got {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

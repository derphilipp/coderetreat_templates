package main //golang-tdd

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	result := add(2, 5)
	expected := 7
	if result != expected {
		t.Fatalf("expected %d, got %d.", expected, result)
	}
}

package main

import (
	"testing"
)

func TestFizzBuzzLength(t *testing.T) {
	fizzbuzz := ReturnFizzBuzz()
	if len(fizzbuzz) != 100 {
		t.Errorf("Length of fizzbuzz was incorrect, got: %d, want: %d.", len(fizzbuzz), 10)
	}
}

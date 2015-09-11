package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	n := 2
	m := 3
	result := Plus(n, m)

	if result != n+m {
		t.Errorf("error plus value: %s", result)
	}
}

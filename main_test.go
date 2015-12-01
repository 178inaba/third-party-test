package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestPlus(t *testing.T) {
	n := 2
	m := 3

	result := Plus(n, m)

	if result != n+m {
		t.Errorf("error plus value: %d", result)
	}
}

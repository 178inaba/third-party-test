package main

import (
	"io"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	so := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Should not be fail: %s.", err)
	}
	os.Stdout = w
	defer func() { os.Stdout = so }()

	if got, want := run(), exitOK; got != want {
		t.Fatalf("ExitStatus is %d, but want %d.", got, want)
	}

	w.Close()
	stdoutBytes, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("Should not be fail: %s.", err)
	}

	if got, want := string(stdoutBytes), "1  +  2  =  3\n"; got != want {
		t.Errorf("Stdout is %s, but want %s.", got, want)
	}
}

func TestPlus(t *testing.T) {
	n := 2
	m := 3
	if got, want := Plus(n, m), n+m; got != want {
		t.Errorf("Plus is %d, but want %d.", got, want)
	}
}

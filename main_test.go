package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRun(t *testing.T) {
	so := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("error os.Pipe():", err)
	}
	os.Stdout = w
	defer func() { os.Stdout = so }()

	run()
	w.Close()

	strCh := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		strCh <- buf.String()
	}()

	result := <-strCh
	answer := "1  +  2  =  3\n"
	if result != answer {
		t.Errorf("error print: %s", result)
	}
}

func TestPlus(t *testing.T) {
	n := 2
	m := 3

	result := Plus(n, m)

	if result != n+m {
		t.Errorf("error plus value: %d", result)
	}
}

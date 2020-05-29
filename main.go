package main

import (
	"fmt"
	"os"
)

const (
	exitOK = iota
)

func main() {
	os.Exit(run())
}

func run() int {
	n := 1
	m := 2
	result := Plus(n, m)

	fmt.Println(n, " + ", m, " = ", result)

	return exitOK
}

// Plus is adding two arguments.
func Plus(n, m int) int {
	return n + m
}

package main

import (
	"fmt"
	"os"
)

const (
	exitOk = iota
)

func main() {
	os.Exit(run())
}

func run() int {
	n := 1
	m := 2
	result := Plus(n, m)

	fmt.Println(n, " + ", m, " = ", result)

	return exitOk
}

// Plus is adding two arguments.
func Plus(n, m int) int {
	return n + m + 1
}

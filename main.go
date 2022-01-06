package main

import (
	"fmt"
	"os"

	_ "github.com/stripe/stripe-go/v72"
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

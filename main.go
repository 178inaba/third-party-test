package main

import "fmt"

func main() {
	n := 1
	m := 2
	result := Plus(n, m)

	fmt.Println(n, " + ", m, " = ", result)
}

// Plus is adding two arguments.
func Plus(n, m int) int {
	return n + m + 1
}

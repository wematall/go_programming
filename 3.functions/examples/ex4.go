package main

import "fmt"

func main() {
	rez, check := add(2, 3)

	fmt.Printf("a + b = %d, a > b = %t\n", rez, check)
}

func add(a, b int) (int, bool) {
	rez := a + b
	check := a > b
	return rez, check
}
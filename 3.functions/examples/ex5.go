package main

import "fmt"

func main() {
	rez, check := add(7, 11)

	fmt.Printf("a + b = %d, a > b = %t\n", rez, check)
}

func add(a, b int) (rez int, check bool) {
	rez = a + b
	check = a > b

	return rez, check
}
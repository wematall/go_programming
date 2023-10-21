package main

import "fmt"

type MyFunctionAdd func(a, b int) int

func sub(c, a, b int, addFunc MyFunctionAdd) int {
	return c - addFunc(a, b)
}

func main() {
	fmt.Println(sub(4, 3, 10, func(a, b int) int {
		return a + b
	}))

	myFunc := func(a, b int) int {
		return a + b
	}

	fmt.Println(sub(4, 3, 10, myFunc))
}
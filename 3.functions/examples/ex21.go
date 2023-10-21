package main

import "fmt"

type MyFunctionAdd func(a, b int) int

func add(a, b int) int {
	return a + b
}

func sub(c, a, b int, addFunc MyFunctionAdd) int {
	return c - addFunc(a, b)
}

func main() {
	myFunc := add
	fmt.Println(sub(4, 3, 10, add))
	fmt.Println(sub(4, 3, 10, myFunc))
}
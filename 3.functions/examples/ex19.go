package main

import "fmt"

func add(a, b int) int{
	return a + b
}

func main() {
	myFunc := add

	fmt.Println(add(3, 5))
	fmt.Println(myFunc(7, 3))
}
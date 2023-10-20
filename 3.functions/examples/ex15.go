package main

import "fmt"

type MyType map[string]map[string][]int

func myFunc(value MyType) {
	fmt.Println(value)
}

func main() {
	myMap := MyType{
		"a":{"a2":[]int{2, 5, 3}},
		"b":{"bbc":[]int{0, 10, 3}},
		"c":{"alex":[]int{}},
	}

	myFunc(myMap)
}
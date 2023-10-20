package main

import "fmt"

func myFunc(value map[string]map[string][]int) {
	fmt.Println(value)
}

func main() {
	myMap := map[string]map[string][]int{
		"a":{"a2":[]int{2, 5, 3}},
		"b":{"bbc":[]int{0, 10, 3}},
		"c":{"alex":[]int{}},
	}

	myFunc(myMap)
}
package main

import "fmt"

var global string = "Global value"

func printValue() {
	fmt.Println(global)
}

func main() {
	fmt.Println(global)
	printValue()

	if true {
		local := 20
		fmt.Println(local)
	}
	// fmt.Println(local) // unavailabel value
	// local = 5 // unavailable value
}
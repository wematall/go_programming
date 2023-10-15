package main

import "fmt"

func print_string(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d: %s\n", i, string(str[i]))
	}
}

func main() {
	str := "some string"

	print_string(str)
}
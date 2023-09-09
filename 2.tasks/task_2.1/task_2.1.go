package main

import (
	"fmt"
)

func main() {
	var s string = "Данная глава была посвящена больше синтаксису Go"
	// s := "hello"
	byteArray := []rune(s)

	for i := 0; i < len(byteArray); i++ {
		if i % 2 == 0 {
			fmt.Println(string(byteArray[i]))
		}
	}
	fmt.Println()
	fmt.Println(s)
}
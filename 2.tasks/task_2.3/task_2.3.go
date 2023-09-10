package main

import "fmt"

func main() {
	s := "Данная глава была посвящена больше синтаксису Go"
	runeSlice := []rune(s)

	for i, v := range runeSlice {
		if i % 6 == 2 || i % 6 == 4 || i % 6 == 5 {
			fmt.Print(string(v))

		}
	}
	fmt.Println()

	fmt.Println(s)
}
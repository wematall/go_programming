package main

import "fmt"

func main() {
	s := "Данная глава была посвящена больше синтаксису Go"
	runeSlice := []rune(s)

	for i, v := range runeSlice {
		if i % 3 == 0 && i % 4 != 0 {
			fmt.Println(string(v))
		}
	}
	fmt.Println(s)
}

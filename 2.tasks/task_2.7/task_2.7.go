package main

import "fmt"

func main() {
	// original string
	s := "список доступных атрибутов"
	var resultSlice []byte

	for i := 0; i < len(s); i++ {
		resultSlice = append(resultSlice, s[i])
	}
	fmt.Println(s)
	fmt.Println(resultSlice)
	fmt.Println(string(resultSlice))
}
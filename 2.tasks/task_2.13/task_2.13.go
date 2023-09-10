package main

import "fmt"

func main() {
	mySlice := make([]int, 0)

	for i := 3; i < 16; i += 4 {
		mySlice = append(mySlice, i)
	}

	fmt.Println(mySlice)
}
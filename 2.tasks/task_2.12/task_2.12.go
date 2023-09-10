package main

import "fmt"

func main() {
	mySlice := make([]int, 0)

	for i := 23; i <= 35; i++ {
		mySlice = append(mySlice, i)
	}
	fmt.Println(mySlice)
}
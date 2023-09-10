package main

import "fmt"

func main() {
	mySlice := []int{3, 0, 1, 3, 0, 4, 3, 4, 56, 6, 1, 3}

	fmt.Println(mySlice)

	for i := len(mySlice)-1; i >= 0; i-- {
		fmt.Printf("%d" + " ", mySlice[i])
	}
}
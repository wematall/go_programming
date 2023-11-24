package main

import "fmt"

func multy(slice []int) int {
	multy := 1

	for _, el := range slice {
		multy *= el
	}

	return multy
}

func main() {
	arr := []int{3, 2, 7}

	fmt.Println(multy(arr))
}
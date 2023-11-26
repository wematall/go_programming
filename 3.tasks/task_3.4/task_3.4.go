package main

import (
	"fmt"
)

// function removes duplicated element in a slice
func removeDuplicates(slice []int) []int {
	slice_tmp := make([]int, 0)
	for i := range slice {
		if !isPresent(slice_tmp, slice[i]) {
			slice_tmp = append(slice_tmp, slice[i])
		}
	}
	return slice_tmp
}

// function check if element present in a slice
func isPresent(slice []int, el int) bool {
	for i := range slice {
		if el == slice[i] {
			return true
		}
	}
	return false
}


func main() {
	arr := []int{7, 3, 5, 7, 3, 11}
	// tmp := make([]int, len(arr))

	fmt.Println("Original slice: ", arr)
	a := removeDuplicates(arr)
	fmt.Println("Remove duplicates: ", a)

}
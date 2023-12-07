package main

import (
	"fmt"
	"slices"
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

// function change element in a slice
// using pointers

func changeElement(array *[]int, a, b int) {
	slices.Replace(*array, a, a+1, b)
}


func main() {
	arr := []int{7, 3, 5, 7, 3, 11}
	// tmp := make([]int, len(arr))

	fmt.Println("Original slice: ", arr)
	a := removeDuplicates(arr)
	fmt.Println("Remove duplicates: ", a)

	fmt.Println(arr)

	changeElement(&arr, 0, 11)
	fmt.Println(arr)

}
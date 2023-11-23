package main

import "fmt"

func sum(slice []int) int {
	sum := 0

	for _, el := range slice {
		sum += el
	}
	return sum
}

func sum_recurcive(slice []int) int {
	sum := 0
	if len(slice) <= 1 {
		return slice[0]
	}

	sum = slice[0] + sum_recurcive(slice[1:])
	return sum
}

func main() {
	slc := []int{3, 7, 5}

	sum := sum(slc)

	fmt.Printf("Sum of slice elements: %d\n", sum)
	fmt.Println(sum_recurcive(slc))
}
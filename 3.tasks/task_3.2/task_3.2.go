package main

import "fmt"

func sum(slice []int) int {
	sum := 0

	for _, el := range slice {
		sum += el
	}
	return sum
}

func main() {
	slc := []int{3, 7, 5}

	sum := sum(slc)

	fmt.Printf("Sum of slice elements: %d\n", sum)
}
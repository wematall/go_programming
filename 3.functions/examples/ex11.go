package main

import "fmt"

func main() {
	array := []float64{2.4, 5.6, 8.1, 9.22}

	fmt.Printf("Array in main: %v\n", array)
	myFunc(&array)
}

func myFunc(array *[]float64) {
	for i := range *array {
		(*array)[i] += 1.5
	}
	fmt.Printf("Array in myFunc(): %v", *array)
}
package main

import "fmt"

func main() {
	array := [4]float64{2.4, 5.6, 8.1, 9.22}
	fmt.Println(array)
	myFunc(&array)

	fmt.Println(array)
}

func myFunc(array *[4]float64) {
	for i := range array {
		array[i] += 2.33
	}
	fmt.Printf("Array in myFunc: %v\n", *array)
}
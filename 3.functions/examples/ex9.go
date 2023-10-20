package main

import "fmt"

func main() {
	a := 7

	fmt.Println(a)

	addSingleValue(a)
	fmt.Println(a)

	array := [4]float64{2.4, 5.6, 8.1, 9.22}
	fmt.Println(array)
	myFunc(array)
	fmt.Println(array)

}

func addSingleValue(a int) {
	a += 11
	fmt.Println(a)
}

func myFunc(array [4]float64) {
	for i := range array {
		array[i] += 2.33
	}
	fmt.Printf("Array in my function: %v\n", array)
}
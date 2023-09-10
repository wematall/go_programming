package main

import "fmt"

func main() {
	myArray := []int{3, 0, 1, 3, 0, 4, 3, 3, 4, 56, 6, 1, 3}
	counter := 0

	for _, v := range myArray {
		if v == 3 {
			counter++
		}
	}
	fmt.Println(myArray)
	fmt.Println(counter)
}
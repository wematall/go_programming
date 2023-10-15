package main

import "fmt"


func swap_items(data []int) {
	for i := 0; i < len(data)-1; i++ {
		data[i], data[i+1] = data[i+1], data[i]
		// fmt.Println(data)
	}
}

func main() {
	data := []int{13, 7, 3, 5, 11}

	fmt.Println(data)
	swap_items(data)

	fmt.Println(data)
}
package main

import "fmt"

// Удалите из среза
// [1, 3, 4, 1, 4, 5, 7, 3, 8, 5]
// повторяющиеся значения

func main() {
	list := []int{1, 3, 4, 1, 4, 5, 7, 3, 8, 5}
	counter := 0

	fmt.Println(list)
	
	for _, v := range list {
		counter = 0
		for k, el := range list {
			if v == el {
				counter++
			}
			if v == el && counter > 1 {
				list = append(list[:k], list[k+1:]...)
			}
		}
	}
	fmt.Println(list)
}
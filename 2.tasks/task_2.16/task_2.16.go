package main

import "fmt"

// выведите различными способами в консоль
// элементы списка
// 3, 0, 1, 3, 0, 4, 3, 4, 56, 6, 1, 3
// с их индексами

func main() {
	list := []int{3, 0, 1, 3, 0, 4, 3, 4, 56, 6, 1, 3}

	// variant 1
	fmt.Println("Вариант 1")
	for i, v := range list {
		fmt.Println(i, v)
	}

	// variant 2
	fmt.Println("Вариант 2")

	for i, v := range list {
		fmt.Printf("%d %d\n", i, v)
	}
}
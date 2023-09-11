package main

// Найдите среднее арифметическое значение элементов
// списка [4, 5, 6, 7, 30, 22, 2, 39, 41]

import "fmt"

func main() {
	result := 0
	sum := 0
	counter := 0

	list := []int{4, 5, 6, 7, 30, 22, 2, 39, 41}

	for _, v := range list {
		sum += v
		counter++
	}
	result = sum / counter
	fmt.Println(result)
}
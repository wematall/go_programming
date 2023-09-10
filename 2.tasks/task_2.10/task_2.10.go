package main

import "fmt"

func main() {
	// считаем сумму чисел в массиве
	summ :=0
	// numbers := []int{3, 0, 1}
	numbers := []int{3, 0, 1, 3, 0, 4, 3, 4, 56, 6, 1, 3}

	for _, v := range numbers {
		summ += v
	}

	fmt.Println(summ)
}
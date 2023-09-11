package main

// Сформируйте срез, значения элементов которого
// находятся в диапазоне от 3 до 25 и без остатка
// делятся на 3

import "fmt"

func main() {
	mySlice := make([]int, 0)

	for i := 3; i < 26; i++ {
		if i % 3 == 0 {
			mySlice = append(mySlice, i)
		}
	}

	fmt.Println(mySlice)
}
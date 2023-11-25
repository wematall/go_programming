package main

import "fmt"

// функция возвращает произведение
// элементов среза
func multy(slice []int) int {
	multy := 1

	for _, el := range slice {
		multy *= el
	}

	return multy
}

// функция возвращает произведение
// элементов среза
// использую рекурсию
func multy_recurcive(slice []int) int {
	// result := 1
	if len(slice) <= 1 {
		return slice[0]
	}
	return slice[0] * multy_recurcive(slice[1:])
}

// using pointers
func multy_pointers(slice *[]int) int {
	result := 1
	for i := range *slice {
		result *= (*slice)[i]
	}

	return result
}

func main() {
	arr := []int{3, 2, 7}

	fmt.Println(multy(arr))
	fmt.Println(multy_recurcive(arr))
	fmt.Println(multy_pointers(&arr))
}
package main

import "fmt"

func min_max(data []int) (int, int) {
	max_el := data[0]
	min_el := data[0]

	for _, value := range data {
		if max_el < value {
			max_el = value
		}
		if min_el > value {
			min_el = value
		}
	}
	return max_el, min_el
}

func main() {
	data := []int{4, 5, 6, 7, 30, 22, 2, 39, 41}

	max_el := data[0]
	min_el := data[0]

	for _, value := range data {
		if max_el < value {
			max_el = value
		}
		if min_el > value {
			min_el = value
		}
	}

	max_result, min_result := min_max(data)

	fmt.Printf("max value: %v, min value: %v\n", max_el, min_el)
	fmt.Printf("max value: %v, min value: %v\n", max_result, min_result)
}
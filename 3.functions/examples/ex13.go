package main

import "fmt"

func is_elementPresent1(slice *[]int, value int, check *bool) {
	*check = false
	for i := range *slice {
		if (*slice)[i] == value {
			*check = true
			break
		}
	}
}

func is_elementPresent2(value int, slice *[]int) bool {
	for i := range *slice {
		if (*slice)[i] == value {
			return true
		}
	}
	return false
}

func main() {
	slice := []int{2, 4, 5, 7, 103, 55}
	var check bool // false by default
	value := 7
	fmt.Println(slice)
	fmt.Println(check)

	is_elementPresent1(&slice, value, &check)
	fmt.Printf("%d contains in slice? %t\n", value, check)
	value = 22
	fmt.Printf("%d contains in slice? %t\n", value, is_elementPresent2(value, &slice))
}


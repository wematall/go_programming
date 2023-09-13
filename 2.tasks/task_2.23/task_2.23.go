package main

// Используя оператор switch-case выведите в терминал
// уведомление о том, какое значение из диапазона от 0 до 5
// подается на его вход

import "fmt"

func main() {
	s := "This is number: "
	var a int
	fmt.Scan(&a)

	switch a {
	case 0:
		fmt.Println(s, 0)
	case 1:
		fmt.Println(s, 1)
	case 2:
		fmt.Println(s, 2)
	case 3:
		fmt.Println(s, 3)
	case 4:
		fmt.Println(s, 4)
	case 5:
		fmt.Println(s, 5)
	default:
		fmt.Println("Unknown value")
	}
}
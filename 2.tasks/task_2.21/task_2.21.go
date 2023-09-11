package main

// Посредством оператора switch-case выведите в терминал
// уведомление является подаваемая на вход буква гласной или согласной

import "fmt"

func main() {
	var result string
	var letter string
	msg1 := "гласная буква"
	msg2 := "эта буква не гласная"

	fmt.Scan(&letter)

	switch letter {
	case "а":
		result = msg1
	case "у":
		result = msg1
	case "о":
		result = msg1
	case "ы":
		result = msg1
	case "э":
		result = msg1
	case "я":
		result = msg1
	case "ю":
		result = msg1
	case "ё":
		result = msg1
	case "и":
		result = msg1
	case "е":
		result = msg1
	default:
		result = msg2

	}
	fmt.Println(result)
}
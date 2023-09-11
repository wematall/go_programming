package main

// Напишите программу, которая считывает целое число
// (месяц), после чего выводит сезон, к которому этот
// месяц относится

import "fmt"

func main() {
	var season string
	var month int
	fmt.Scan(&month)

	if month == 12 || month == 1 || month == 2 {
		season = "Winter"
	} else if month >= 3 && month <= 5 {
		season = "Spring"
	} else if month >= 6 && month <= 8 {
		season = "Summer"
	} else if month >= 9 && month <= 11 {
		season = "Autumn"
	} else {
		season = "unknown"
	}

	fmt.Println(season)
}
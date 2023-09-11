package main

// Напишите программу, выводящую среднее из трех значений

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (a < b && a > c) || (a < c && a > b){
		fmt.Println(a)
	} else if (b > c && b < a) || (b > a && b < c) {
		fmt.Println(b)
	} else if (c > a && c < b) || (c > b && c < a){
		fmt.Println(c)
	} else {
		fmt.Println("непонятно")
	}
}
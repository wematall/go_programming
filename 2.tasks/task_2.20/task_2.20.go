package main

// Напишите программу, выводящую таблицу умножения
// для задаваемого пользователем числа
// от 1 до 9 включительно

import "fmt"

func main() {

	var a int
	var result int

	fmt.Println("Печатаем таблицу умножения")
	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	fmt.Println()

	for i := 1; i < 10; i++ {
		result = a * i
		fmt.Printf("%d * %d = %d\n", a, i, result)
	}
}
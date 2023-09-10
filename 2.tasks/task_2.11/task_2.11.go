package main

import "fmt"

func main() {
	list := []int{3, 0, 1, 3, 0, 4, 3, 4, 56, 6, 1, 3}
	summ := 0

	for i, v := range list {
		if i % 3 == 0 {
			summ += v
		}
	}
	fmt.Println("Summa: ", summ)
}
package main

import "fmt"

func main() {
	sum := func(values ...int) (sum int) {
		for _, v := range values {
			sum += v
		}
		return
	}(3, 5, 9, 25, -10)

	fmt.Printf("Sum = %d\n", sum)
}
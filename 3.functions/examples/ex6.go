package main

import "fmt"

func main() {
	res := sum(3, 5, 9, 25, -10)
	fmt.Printf(" Sum = %d\n", res)
}

func sum(values ...int) int{
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
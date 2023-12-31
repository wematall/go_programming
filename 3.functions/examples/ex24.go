package main

import (
	"fmt"
	"math"
)

type MyPow func(value int) int

func degree(degree int) MyPow {
	return func(value int) int {
		return int(math.Pow(float64(value), float64(degree)))
	}
}
func main() {
	calculation := degree(3)
	fmt.Println(calculation(3))
	fmt.Println(calculation(4))

	calculation = degree(9)
	fmt.Println(calculation(4))
	fmt.Println(calculation(7))
}
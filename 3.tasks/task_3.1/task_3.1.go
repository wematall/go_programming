package main

import "fmt"

func max_abc(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}


func main() {
	result := max_abc(7, 3, 11)
	fmt.Printf("Max: %d\n", result)
}
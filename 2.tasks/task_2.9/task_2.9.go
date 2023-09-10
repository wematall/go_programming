package main

import "fmt"

func main() {
	// print numbers
	// 5 and 7 don't print
	for i := 1; i < 10; i++ {
		if i == 5 || i == 7 {
			continue
		} else {
			fmt.Printf("%d" + " ", i)
		}
	}
}
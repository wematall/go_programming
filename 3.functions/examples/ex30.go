package main

import "fmt"

func main() {
	defer fmt.Printf("%d\n", 0)
	defer fmt.Printf("%d\n", 1)
	defer fmt.Printf("%d\n", 2)
	fmt.Println("Finish!")
}
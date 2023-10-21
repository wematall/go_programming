package main

import "fmt"

var global string = "Global value"

func printVal() {
	fmt.Println(global)
}


func main() {
	fmt.Println(global)
	printVal()
}
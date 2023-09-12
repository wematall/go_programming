package main

import "fmt"

func main() {
	// объявляем по разному
	map1 := make(map[string]int)

	map2 := map[string]int{
		"a":1,
		"b":2,
	}

	var map3 map[string]int

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
}
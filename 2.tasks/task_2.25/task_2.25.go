package main

import "fmt"

func is_keyPresent(str string, data map[string]int) bool {
	_, ok := data[str]
	
	if ok {
		return true
	} else {
		return false
	}
}

func main() {
	data := make(map[string]int)
	data["apple"] = 3
	data["banana"] = 5
	data["onion"] = 7

	result := is_keyPresent("ananas", data)
	fmt.Println(result)

	fmt.Println(is_keyPresent("banana", data))
}
package main

import "fmt"

func main() {
	// str := "google.com"
	str := "hello golang"
	elements := []rune(str)

	myMap := make(map[rune]int)

	for _, value := range elements {
		myMap[value]++

	}
	fmt.Println(myMap)
	fmt.Println(str)


	for key, value := range myMap {
		fmt.Printf("%c: %d\n", key, value)
	}

}
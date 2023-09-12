package main

import "fmt"

func main() {
	mapa := make(map[string]int)
	mapa["apple"] = 10
	mapa["banana"] = 12
	mapa["ananas"] = 33

	element_0, ok_0 := mapa["onion"] // element is not present
	element_ananas, ok_ananas := mapa["ananas"]


	fmt.Println(mapa)
	fmt.Println("element: ", element_0)
	fmt.Println(ok_0)
	fmt.Println("element_ananas", element_ananas)
	fmt.Println(ok_ananas)
}
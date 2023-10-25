package main

import "fmt"

type microwaveCreator func(dish string, model int) string

var microwaveIndex int = 0

func factory(microwaveName string, power int) microwaveCreator {
	microwaveIndex++
	model := fmt.Sprintf("%s-RU-001%d", microwaveName, microwaveIndex)

	return func(dish string, mode int) string {
		str := fmt.Sprintf("Микроволновка %s мощность %d w Вт,", model, power)
		str += fmt.Sprintf("Греет блюдо %s в режиме %d", dish, mode)
		
		return str
	}
}


func main() {
	microwave := factory("Scarlet", 800)
	fmt.Println(microwave("Суп", 3))
	fmt.Println(microwave("Пюре", 2))
	newMicrowave := factory("LG", 1200)
	fmt.Println(newMicrowave("Плов", 4))
}

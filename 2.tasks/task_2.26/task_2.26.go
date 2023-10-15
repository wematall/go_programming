package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	data := []int{15, 23, 150}
	var str string

	for _, value := range data {
		str = str + string(strconv.Itoa(value))
	}

	result, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result)) // check type of result
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func sliceCreator1(size int16) (slice *[]int16) {
	slice = new([]int16)
	*slice = make([]int16, size, size+10)

	for i := range *slice {
		(*slice)[i] += int16(generator.Int())
	}
	return
}

func sliceCreator2(size int16) *[]int16 {
	slice := make([]int16, size, size+10)
	for i := range slice {
		slice[i] += int16(generator.Int())
	}
	return &slice
}

func main() {
	fmt.Println(int16(generator.Int()))

	slice := sliceCreator1(7)
	fmt.Println(*slice)

	slice = sliceCreator2(5)
	fmt.Println(*slice)
}
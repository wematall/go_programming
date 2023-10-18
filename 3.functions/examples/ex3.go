package main

import "fmt"

func main() {
	fmt.Println(userInfo(3, 5, "Alex"))
}

func userInfo(a, b int, name string) string {
	return fmt.Sprintf("user name: %s, age: %d, weight: %d\n", name, a, b)
}
package main

import "fmt"

var global string = "Global value"

func main() {
	fmt.Println(global)

	local := 20
	{
		fmt.Println(local)
		{
			local := 10
			fmt.Println(local)
		}
		fmt.Println(local)
	}
}
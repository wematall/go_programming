package main

import "fmt"

// собрать две новых строки из исходной по индексам четные нечетные
func main() {
	str := "some long string"

	var str1, str2 string;

	for i := 0; i < len(str); i++ {
		if i % 2 == 0 {
			str1 = str1 + string(str[i])
		} else {
			str2 = str2 + string(str[i])
		}
	}

	fmt.Println(str)	

	fmt.Println(str1)
	fmt.Println(str2)

}
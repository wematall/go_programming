package main

import (
	"fmt"
	"sort"
)

func sort_mapa(data map[string]int) {
		// отсортируем
	keys := make([]string, 0, len(data))
 
    for key := range data {
        keys = append(keys, key)
    }
 
    sort.SliceStable(keys, func(i, j int) bool{
        return data[keys[i]] > data[keys[j]]
    })
		i := 0
		for _, k := range keys{
			fmt.Printf("%s: %d\n", k, data[k])
			i++
			if i == 1 {
				break
			}
		}
}


func main() {
	str := "some long string data banana"
	data_elements := make(map[string]int)

	for _, el := range str {
		data_elements[string(el)]++
	}

	fmt.Println("sorted result")

	sort_mapa(data_elements)

}
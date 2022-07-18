// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
)

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	map_str := make(map[string]string)
	var uniqueSet []string

	for _, v := range str {
		if _, ok := map_str[v]; !ok {
			map_str[v] = ""
		}
	}
	for key, _ := range map_str {
		uniqueSet = append(uniqueSet, key)
	} 

	fmt.Println("Исходная строка", str)
	fmt.Println("Cобственное множество", uniqueSet)
}
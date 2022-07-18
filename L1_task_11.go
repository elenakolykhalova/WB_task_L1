//Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func createMap(a []string, b map[string]int) {
	for _, v := range a {
		if _, ok := b[v]; !ok {
			b[v] = 1
		} else {
			b[v] = 2
		}
	}
}

func intersectionOFsets(a, b []string) []string {
	map_sets := make(map[string]int)
	createMap(a, map_sets)
	createMap(b, map_sets)
	var c []string
	for key, _ := range map_sets {
		if map_sets[key] == 2 {
			c = append(c, key)
		}
	}
	return c
}


func main() {
	a := []string{"cat", "dog", "luna", "jhone"}
	b := []string{ "cat", "luna", "kate", "jhone"}

	fmt.Println(intersectionOFsets(a, b))
}
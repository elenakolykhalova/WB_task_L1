// Разработать программу, которая переворачивает подаваемую на ход строку 
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func main() {
	var str, res string
	fmt.Scan(&str)

	for _, i := range str {
		res = string(i) + res
	}
	fmt.Println(res)
}
// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func ReplaceNBit(num, pos, value int64) int64 {
	if value >= 1{
		return num | (int64(1) << pos) // побитовое или
	}
	return num &^ (int64(1) << pos) // и не
}

func main(){
	var a int64 = 7
	fmt.Printf("Before: %08b\n", a)
	fmt.Printf("After : %08b\n", ReplaceNBit(a, 1, 0))
}
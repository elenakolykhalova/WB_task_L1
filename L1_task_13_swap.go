// Поменять местами два числа без создания временной переменной.

package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 67

	a, b = b, a
	fmt.Printf("Значение а = %d, b = %d\n", a, b)
}

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var first, second string
	var a, b big.Int

	fmt.Scan(&first, &second)

	_, a_ok := a.SetString(first, 10)
	_, b_ok := b.SetString(second, 10)
	if !a_ok || !b_ok {
		fmt.Println("Not digit")
	} else {
		mul := big.NewInt(0).Mul(&a, &b)
		div := big.NewInt(0).Div(&a, &b)
		sum := big.NewInt(0).Add(&a, &b)
		diff := big.NewInt(0).Sub(&a, &b)
	
		fmt.Printf("x * y = %d\n", mul)
		fmt.Printf("x / y = %d\n", div)
		fmt.Printf("x + y = %d\n", sum)
		fmt.Printf("x - y = %d\n", diff)
	}
}
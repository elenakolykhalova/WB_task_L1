// Удалить i-ый элемент из слайса.

package main

import (
	"fmt"
)

// //без сохранения порядка
// func main() {
// 	sliceExample := []int {45, 55, 66, 59, 69, 100}
// 	i := 3
// 	// Перезаписываем на i-ый элемент последний элемент слайса
// 	sliceExample[i] = sliceExample[len(sliceExample)-1]
// 	// Делаем усечение
// 	sliceExample = sliceExample[:len(sliceExample)-1]
// 	fmt.Println(sliceExample)
// }

//с сохранением порядка
func main() {
	sliceExample := []int {45, 55, 66, 3, 69, 100}
	i := 3
	//делаю сдвиг влево на один элемент
	copy(sliceExample[i:], sliceExample[i+1:])
	//делаю усечение
	sliceExample = sliceExample[:len(sliceExample)-1]
	fmt.Println(sliceExample)
}
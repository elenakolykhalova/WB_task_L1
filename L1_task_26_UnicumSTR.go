// Разработать программу, которая проверяет, что все символы в строке уникальные 
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// 	abcd — true 
// 	abCdefAaf — false 
// 	aabcd — false

package main

import (
	"strings"
	"fmt"
)

func main() {
	var str string
	
	fmt.Scan(&str)
	//сделали все буквы маленькими
	str = strings.ToLower(str)
	//напечатали результат функции мапчек
	fmt.Println(Mapcheck(str))
}


func Mapcheck(str string) bool {
	//безопасный способ создания мапы
	abc := make(map[rune]int)
	//начинаем итерироваться по строке
	for _, i := range str {
		//смотрим, есть ли в мапе ключ i, если есть, возвращаем false, 
		// иначе добавляем в мапу новый ключ
		if _, ok := abc[i]; !ok {
			abc[i] = 1
		} else {
			return false
		}
	}
	return true
}
	
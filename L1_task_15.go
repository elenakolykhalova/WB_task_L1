// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

package main

import "fmt"

var justString string 

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 { //потому включаем проверку на длину строки v
		justString = v[:100] // опасность, что самописная функция createHugeString вернет строку меньше 100 символов
		fmt.Println(justString)
	} else {
		fmt.Println("Недостаточная длина строки")
	}
}

func main() { 
	someFunc()
}

func createHugeString(i int) string {
	return "At the time, no single team member knew Go, but within a month, everyone was writing in Go and we we"
}



// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//записываем все слова из строки ввода в буфер до '\n'
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err) 
	}
	//удаляем '\n' в конце строки
	text = strings.Trim(text, "\n") 
	// делаем слайс разделяя строку по пробелам
	sliceWord := strings.Split(text, " ")
	//выводим слайс, начиная с последнего элемента
	i := len(sliceWord) - 1
	for i >= 0 {
		fmt.Print(sliceWord[i], " ")
		i--
	}
	fmt.Println()
}
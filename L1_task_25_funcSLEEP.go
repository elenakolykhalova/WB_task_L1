// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func mySleep(msDuration int) {
	end := time.Now().Add(time.Duration(msDuration * int(time.Millisecond))) // когда мы должны проснуться
	fmt.Println("Я засыпаю...")
	for time.Now().Before(end){ // запускаем пустой цикл while time.Now() для отсчета времени
	}
	fmt.Println("Я проснулся...")
}

func main(){
	fmt.Println(time.Now())
	mySleep(5000)
	fmt.Println(time.Now())
}

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, 
// во второй — результат операции x*2, 
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	A := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("Исходный массив", A)
	firstChan := make(chan int, len(A)) //создаем буферизированный канал в размер длины массива
	secondChan := make(chan int, len(A))
	wg.Add(3)
	go func(){
		for _, v := range A {
			firstChan <- v
		}
		wg.Done()
		close(firstChan)
	}()
	go func(){
		for valueFromFirstChan := range firstChan {
			secondChan <- valueFromFirstChan * valueFromFirstChan
		}
		wg.Done()
		close(secondChan)
	}()
	go func(){
		for valueFromSecondChan := range secondChan {
			fmt.Println("Значениe из второго канала", valueFromSecondChan)
		}		
		wg.Done()
	}()
	wg.Wait()
}
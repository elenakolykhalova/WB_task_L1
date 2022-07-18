// Реализовать все возможные способы остановки выполнения горутины.


package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Способы остановки выполнения горутины
// Из пакета contex:
// 1.WithTimeOut
// 2.WithDeadline
// 3.WithCancel
// И Ручная остановка горутины

func WithCancel(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
			case <- ctx.Done():
				fmt.Println("Остановка через cancel")
				wg.Done()
				return
			default:
				fmt.Println(i)
				time.Sleep(time.Second)
		}
	}
}

func WithTimeOut(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
			case <- ctx.Done():
				fmt.Println("Остановка через timeOut")
				wg.Done()
				return
			default:
				fmt.Println(i)
				time.Sleep(time.Second)
		}
	}
}

func WithDeadline(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i >= 0; i++ {
		select {
			case <- ctx.Done():
				fmt.Println("остановка через Deadline")
				wg.Done()
				return
			default:
				fmt.Println(i)
				time.Sleep(time.Second)
			}
	}
}

func HandMadeCancelChan(cancel chan bool, wg *sync.WaitGroup) {
	for i := 0; i >= 0; i++ {
		select {
			case <- cancel:
				fmt.Println("Ручная остановка")
				wg.Done()
				return
			default:
				fmt.Println(i)
				time.Sleep(time.Second)
			}
	}
}

func main(){
	var wg sync.WaitGroup
	mainCtx := context.Background()
	cancelChan := make(chan bool) // Создаем канал для ручной остановки канала

	wg.Add(4)
	ctx1, cancel1 := context.WithCancel(mainCtx) // Когда вызывается cancel1() останавливается горутина <-ctx.Done
	ctx2, _ := context.WithTimeout(mainCtx, 2 * time.Second) // получает n секунд до вызова cancel()
	ctx3, _ := context.WithDeadline(mainCtx, time.Now().Add(10 * time.Second)) // получает метку времени, когда будет вызвана функция cancel()
	go WithCancel(ctx1, &wg)
	go WithDeadline(ctx3, &wg)
	go WithTimeOut(ctx2, &wg)
	go HandMadeCancelChan(cancelChan, &wg)
	go func(){
		time.Sleep(6 * time.Second) // выжидаем 6 секунд ctx cancel managment
		cancel1() //вызываем cancel1 для остановки горутины через ctx.Done
		time.Sleep(1 * time.Second)
		cancelChan <- true // канал для ручной остановки заполнен => сигнал для остановки
	}()
	wg.Wait()
}
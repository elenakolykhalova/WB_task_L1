// Реализовать постоянную запись данных в канал (главный поток). 
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. 
// Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func mainWriter(ctx context.Context, wg *sync.WaitGroup, intMun chan int) {
	for i := 0; i >= 0; i++ {
		select{
		case <-ctx.Done(): //  когда вызывается cancel() мы завершаем выполнение горутины
			fmt.Println("Я главный поток и я остановился")
			wg.Done()
			return
		default:
			fmt.Printf("Запись в главный поток: %d\n", i)
			intMun <- i
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func regularWorker(ctx context.Context, wg *sync.WaitGroup, intNum chan int, chanName int,) {
	for i := 0; i >= 0; i++ {
		select{
		case <- ctx.Done(): // когда вызывается cancel() мы завершаем выполнение горутины
			fmt.Println("Я Вокер", chanName, "и я остановился")
			wg.Done()
			return
		case v := <- intNum:
			fmt.Printf("Я Вокер под № %d и я читаю из главного канала: %d\n", chanName, v)
			time.Sleep(time.Millisecond * 4000)
		}
	}
}

func main(){
	var (
		intNumbers chan int
		nWorkers int
		wg	sync.WaitGroup
	)

	ctx, cancel := context.WithCancel(context.Background()) 
	intNumbers = make(chan int)
	fmt.Println("Введите количество вокеров")
	fmt.Scan(&nWorkers)
	wg.Add(1)
	go mainWriter(ctx, &wg, intNumbers) //запускаем запись в главный канал
	//запускаем анонимную функцию для ожидания Ctrl+C
	go func() {
		c := make(chan os.Signal, 1) 
		signal.Notify(c, os.Interrupt)
		<-c // ожидание Ctrl+C  
		cancel() // Если нам нужно послать сигнал, чтобы остановить что-либо, то вызываем cancel().
	}()
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go regularWorker(ctx, &wg, intNumbers, i + 1)
	}
	wg.Wait()
	fmt.Println("Вы нажали CTRL+C")
}

//Разработать программу, которая будет последовательно отправлять значения в канал, 
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sender(ctx context.Context, wg *sync.WaitGroup, putHere chan int){
	for i := 1; i >= 1; i++ {
		select{
			case <- ctx.Done():
				fmt.Println("Отправитель закончил работу")
				wg.Done()
				return
			default:
				putHere <- i
				time.Sleep(time.Second * 1)
		}
	}
}

func reader(ctx context.Context, wg *sync.WaitGroup, getHere chan int) {
	for {
		select {
			case <- ctx.Done():
					fmt.Println("Читатель закончил работу")
					wg.Done()
					return
			case v := <- getHere:
				fmt.Println(v)
		}
	}
}

func main(){
	pipe := make(chan int)
	nSecond := 0
	var wg sync.WaitGroup

	fmt.Println("Через сколько секунд закончить работу?")
	fmt.Scan(&nSecond)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * time.Duration(nSecond))
	wg.Add(2)
	go reader(ctx, &wg, pipe)
	go sender(ctx, &wg, pipe)
	wg.Wait()
}
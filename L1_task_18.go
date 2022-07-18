// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.


//Можно решить двумя путями, через sync.Mutex и через sync/atomic

package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

type count struct {
	i int
}

func main(){
	N := 1574867
	wg := new(sync.WaitGroup)
	mx := new(sync.Mutex)
	var test count

	wg.Add(N)
	for k := 0; k < N; k++ {
		go func(C *count){
			defer wg.Done()
			mx.Lock()
			C.i++
			mx.Unlock()
		}(&test)
	}
	wg.Wait()
	fmt.Println(test.i)
}


// type counter struct {
// 	i int64
// }

// func (c *counter) add() {
// 	atomic.AddInt64(&c.i, 1)
// }

// func main() {
// 	N := 45789
// 	var wg sync.WaitGroup
// 	cnt := counter{i: 0}
	
// 	wg.Add(N)
// 	for j := 0; j < N; j++ {
// 		go func() {
// 			cnt.add()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(cnt.i)
// }
// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex
	myMap := make(map[int] int)
	var counter int
	nWorkers := 100

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go func(myMap map[int] int, key *int){
			defer wg.Done()
			mux.Lock()
			myMap[*key] = rand.Intn(150)
			*key += 1
			mux.Unlock()
		}(myMap, &counter)
	}
	wg.Wait()
	fmt.Println("nWoker =", nWorkers, "\nlen(myMap)", len(myMap))
}
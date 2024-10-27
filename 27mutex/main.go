package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race COndi")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var scores = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("1 func")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("2 func")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3 func")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	mut.RLock()	
	fmt.Println(scores)
	mut.RUnlock()
}

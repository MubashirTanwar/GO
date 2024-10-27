package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")
	ch := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChannelOpen := <-ch

		fmt.Println(val)
		fmt.Println(isChannelOpen)


		fmt.Println(<-ch)
		// fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup){
		ch <- 2
		ch <- 6
		wg.Done()
	}(ch, wg)


	wg.Wait()
}
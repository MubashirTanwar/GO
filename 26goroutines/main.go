package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {
	fmt.Println("GG")
	websites := []string{
		"https://go.dev",
		"https://google.com",
		"https://mubashirtanwar.com",
		"https://go.dev",
	}
	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	// getStatusCode("http://google.com")
	fmt.Println(signals)
}

func getStatusCode(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals = append(signals, url)
	mut.Unlock()
	fmt.Printf("%v status code for %v\n", res.StatusCode, url)
}

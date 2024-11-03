package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"sync"
)

var ch = make(chan string)
var wg = &sync.WaitGroup{}

func whatsTheWeather(city string, ch chan <- string, wg *sync.WaitGroup ) {
	defer wg.Done()
	start := time.Now()
	// ... operation that takes 20 milliseconds ...
	var url = "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=85fdad722cfc9635e004eb08471d2ab2"
	_, err := http.Get(url)
	checkNilErr(err)

	t := time.Now()
	elapsed := t.Sub(start)
	ch <- fmt.Sprintf("%s time: %s", city, elapsed)
}

func main() {
	var cities = []string{
		"Mumbai", "Delhi", "Jaipur", "Sikar",
	}
	start := time.Now()
	for _, city := range cities {
		wg.Add(1)
		go whatsTheWeather(city, ch, wg)
	}
	go func(){
		wg.Wait()
		close(ch)
	}()
	for result := range ch {
		fmt.Println(result)
	}
	end := time.Now()
	fmt.Printf("Total time taken %v", end.Sub(start))
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

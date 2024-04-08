package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	present := time.Now()
	fmt.Println(present)

	fmt.Println("Current Month",present.Format("01"))

	createdDate := time.Date(2020, time.April, 24, 12, 00, 00, 00, time.UTC)
	fmt.Println(createdDate.Format("Monday"))
}

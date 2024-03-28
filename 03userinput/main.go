package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomr := "Welcome User: "
	fmt.Println(welcomr)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")

	input, _ := reader.ReadBytes(32) 
	// input, _ := reader.ReadByte(32)  this reads a single byte
	// WOW
	// 32 means space
	// 13 means enter
	// will read bytes until it hits the given unicode
	fmt.Println("Your Age is ", input)
	fmt.Printf("Your Type is %T", input)
}

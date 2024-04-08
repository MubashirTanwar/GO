package main

import "fmt"

func main() {
	fmt.Println("We GO Array")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"

	fmt.Println(len(fruits))
}
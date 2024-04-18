package main

import (
	"fmt"
)

func main() {
	fmt.Println("We GO Array")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"

	var nums = [6]int{}
	nums[1] = 234
	nums[2] = 3456
	nums[5] = 56658

	fmt.Printf("%T", nums)
	fmt.Printf("%T", fruits)
}

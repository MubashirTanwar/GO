package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("we are doing switch")
	var diceNumber int = rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println(`Move ONE step`)

	case 2:
		fmt.Println(`Move TWO step`)

	case 3:
		fmt.Println(`Move THREE step`)
		fallthrough
	case 4:
		fmt.Println(`Move FOUR step`)
	case 5:
		fmt.Println(`Move FOUR step`)
	case 6:
		fmt.Println(`Move FOUR step`)
	default:
	    fmt.Println("NIGGA WHAT EVEN")

	}
}

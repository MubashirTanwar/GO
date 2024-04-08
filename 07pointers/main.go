package main

import "fmt"

func main() {
	fmt.Println("We GO Pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer is", ptr)
	myNum := 23

	var ptrnum *int = &myNum

	fmt.Println(&myNum)
	fmt.Println(*ptrnum)

}

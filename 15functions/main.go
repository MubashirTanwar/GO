package main

import "fmt"

func main() {
	fmt.Println("WE DO FUCNTIONS")
	greet()
	value, message:= proAdded(3,4,5,6,24,54,5)
	fmt.Printf("%v %v", message, value)
	
}

func greet(){
	fmt.Println("THIS IS THE GREET FUNCTION")
}

// func add(intOne int, intTwo int)int{
// 	return intOne + intTwo
// }

func proAdded(values ...int) (int, string) {
	total := 0
	for _, value := range values{
		total += value
	}
	return total, "THE TOTAL IS"
}
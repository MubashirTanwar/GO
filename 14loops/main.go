package main

import(
	"fmt"
)


func main(){
	fmt.Println("WE DO LOOPS")

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d:=0; d<len(days); d++ {
	// 	fmt.Printf(days[d])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days{
	// 	fmt.Printf("day is %v\n", day)
	// }

	var value int = 0

	for value < 10{
		if value == 2 {
			goto test
		}
		// if i == 5 {break}
		fmt.Println(value)
		value++

	}

	test:
	fmt.Printf("A GOTO %v", value)
}
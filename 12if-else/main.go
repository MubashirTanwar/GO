package main

import (
	"fmt"
)

func main() {
	fmt.Println("If Else tutorial")

	// reader := bufio.NewReader(os.Stdin)
	// var joke, _ = reader.ReadString(13)
	// joke = strings.TrimSpace(joke)
	// fmt.Printf("type of joke outside is %T\n", joke)
	// fmt.Printf("your Input is %v\n", joke)

	// if joke == "unfunny" {
	// 	fmt.Printf("type of joke is %T\n", joke)
	// 	fmt.Println("Very unfunny")
	// } else {
	// 	fmt.Printf("type of joke is %T\n", joke)
	// 	fmt.Println("Funny")
	// }

	if num := 4; num < 10 {
		fmt.Println("Less than 10")
	}
}

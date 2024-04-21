package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("If Else tutorial")

	reader := bufio.NewReader(os.Stdin)
	var joke, _  =  reader.ReadString(13)
	fmt.Printf("your Input is %v\n", joke)
	
	if (joke == "unfunny") {
		fmt.Println("Very unfunny")
	} else {
		fmt.Println("Funny")
	}
}

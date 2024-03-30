package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop")
	fmt.Println("Rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)
	input ,_ := reader.ReadString('\n') 
	fmt.Print("your rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your rating is: ", numRating + 1)
	}

}

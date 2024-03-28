package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func main() {
	fmt.Println(reverse.String(quote.Go()))
}

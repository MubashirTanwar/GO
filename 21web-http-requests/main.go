package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("HTTP VERBS")
	GetRequest()
}

func GetRequest() {
	const url = "http://localhost:3333"

	res, err := http.Get(url)
	checkNilError(err)
	defer res.Body.Close()

	fmt.Println("Status Code is", res.StatusCode)
	
	fmt.Println("content lenght", res.ContentLength)
	
	content, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(content))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

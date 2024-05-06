package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://karpathy.ai/"

func main() {
	fmt.Println("WEB REQUESTS")
	res, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Type of response is %T \n", res)
	fmt.Printf("Response %+v \n", res)

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	checkNilErr(err)
	content := string(data)
	fmt.Println("Content Extracted", content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

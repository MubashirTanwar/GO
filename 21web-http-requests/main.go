package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("HTTP VERBS")
	// GetRequest()
	// PostRequest()
	FormRequest()
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

func PostRequest() {

	const url = "http://localhost:3333/post"

	//JSON Payload

	reqBody := strings.NewReader(`
	{
		"coursename":"GoLang",
		"price": 30,
		"platform": "Youtube"
	}
	`)

	res, err := http.Post(url, "application/json", reqBody)
	checkNilError(err)
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(content))
}

func FormRequest() {
	const myUrl = "http://localhost:3333/postform"

	data := url.Values{}
	data.Add("firstName", "Mubashir")
	data.Add("lastName", "Tanwar")
	data.Add("Age", "21")

	res, err := http.PostForm(myUrl, data)
	
	checkNilError(err)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println("Response", string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

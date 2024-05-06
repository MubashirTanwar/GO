package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://trash.ai:300/home/product?id=333&color=blue"

func main() {
	fmt.Println("URL handling")
	res, err := url.Parse(myUrl)
	checkNilErr(err)
	fmt.Println(res.Scheme)
	fmt.Println(res.Path)
	fmt.Println(res.Query())
	fmt.Println(res.RawQuery)
	fmt.Println(res.Port())
	fmt.Println(res.String())
	fmt.Println(res.RawFragment)
	
	qparam := res.Query()
	fmt.Printf("%+v\n", qparam)
	fmt.Println(qparam["id"])

	for _, val := range qparam{
		fmt.Println("Param is", val[0])
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

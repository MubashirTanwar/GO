package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON EXES")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	courses := []course{
		{"GO", 299, "Youtube", "test", []string{"GO", "WEB", "JS"}},
		{"GO", 199, "LCO", "123", []string{"Google", "YT", "HC"}},
		{"GO", 999, "Udemy", "abc", []string{"Udemy", "DEV", "JS"}},
		{"GO", 99, "FCC", "xyz", nil}}

	Json, err := json.MarshalIndent(courses, "", "  ")
	checkNilError(err)

	fmt.Println("JSON", string(Json))
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "GO",
		"Price": 299,
		"website": "Youtube",
		"Tags": [
		  "GO",
		  "WEB",
		  "JS"
		]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON Valid")
		json.Unmarshal(jsonData, &lcoCourse)
		// fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// when you wan tto add data to key value pair

	var OnlineData map[string]interface{}	
	json.Unmarshal(jsonData, &OnlineData)
	fmt.Printf("%#v\n", OnlineData)

	for k, v := range OnlineData {
		fmt.Printf("Key is %v and value is %v of type %T \n",k,v,v)
	}
	fmt.Printf("Type of %T", OnlineData)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
